package dao

import (
	"context"
	"os"
	"time"

	"github.com/XMUMY/lost_found/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() *mongo.Client {
	mongoOptions := options.Client().ApplyURI(os.Getenv("DB_ADDR"))
	mongoOptions.Auth = &options.Credential{
		AuthSource: os.Getenv("DB_NAME"),
		Username:   os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
	}

	// Connect to mongodb.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		panic(err)
	}

	return client
}

// InsertItem add new item to database and returns its id.
func (d *Dao) InsertItem(ctx context.Context, item *model.LostAndFoundDetail) (id string, err error) {
	col := d.db.Collection("items")
	res, err := col.InsertOne(ctx, &item)
	if err != nil {
		err = errors.Wrapf(err, "failed to insert item %+v", item)
		return
	}

	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

// DeleteItemWithUID remove item with specified ID AND UID.
func (d *Dao) DeleteItemWithUID(ctx context.Context, id string, uid string) error {
	col := d.db.Collection("items")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.InvalidItemIDError
	}

	_, err = col.DeleteOne(ctx, bson.M{"_id": objID, "uid": uid})

	return errors.Wrapf(err, "failed to delete %s", id)
}

// FindItem get LostAndFoundItem by id.
func (d *Dao) FindItem(ctx context.Context, id string) (item *model.LostAndFoundItem, err error) {
	col := d.db.Collection("items")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err = model.InvalidItemIDError
		return
	}

	item = new(model.LostAndFoundItem)
	err = col.FindOne(ctx, bson.M{"_id": objID}).Decode(&item)
	if err != nil && err == mongo.ErrNoDocuments {
		err = model.ItemNotFoundError
		return
	}

	err = errors.Wrapf(err, "failed to find item %s", id)
	return
}

// FindItemBriefs get a list of item briefs before (not include) given date.
func (d *Dao) FindItemBriefs(ctx context.Context, date *time.Time) (items []*model.LostAndFoundItem, err error) {
	col := d.db.Collection("items")

	limit := int64(10)
	findOptions := options.FindOptions{
		Limit: &limit,
		Sort:  bson.M{"timestamp": -1},
		// Ignore detail part.
		Projection: bson.M{
			"description": false,
			"contacts":    false,
			"pictures":    false,
		},
	}

	cursor, err := col.Find(ctx, &bson.M{"timestamp": bson.M{"$lt": date}}, &findOptions)
	if err != nil {
		goto end
	}

	for cursor.Next(ctx) {
		result := new(model.LostAndFoundItem)
		err = cursor.Decode(result)
		items = append(items, result)
	}

end:
	err = errors.Wrapf(err, "failed to get briefs before %v", date)
	return
}
