package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/XMUMY/lost_found/api/lost_found/v4"
	"github.com/XMUMY/lost_found/internal/biz"
)

type itemRepo struct {
	data   *Data
	logger *log.Helper
}

func NewItemRepo(data *Data, logger log.Logger) biz.ItemRepo {
	return &itemRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// InsertItem add new item to database and returns its id.
func (r *itemRepo) InsertItem(ctx context.Context, item *biz.ItemDetail) (id string, err error) {
	col := r.data.db.Collection("items")
	res, err := col.InsertOne(ctx, &item)
	if err != nil {
		err = errors.Wrapf(err, "failed to insert item %+v", item)
		return
	}

	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

// FindItem get LostAndFoundItem by id.
func (r *itemRepo) FindItem(ctx context.Context, id string) (item *biz.Item, err error) {
	col := r.data.db.Collection("items")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err = pb.InvalidItemIDError
		return
	}

	item = &biz.Item{}
	err = col.FindOne(ctx, bson.M{"_id": objID}).Decode(&item)
	if err != nil && err == mongo.ErrNoDocuments {
		err = pb.ItemNotFoundError
		return
	}

	err = errors.Wrapf(err, "failed to find item %s", id)
	return
}

// FindItemBriefs get a list of item briefs before (not include) given date.
func (r *itemRepo) FindItemBriefs(ctx context.Context, date *time.Time) (items []*biz.Item, err error) {
	col := r.data.db.Collection("items")

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
		result := &biz.Item{}
		err = cursor.Decode(result)
		items = append(items, result)
	}

end:
	err = errors.Wrapf(err, "failed to get briefs before %v", date)
	return
}

// DeleteItemForUser remove item with specified ID and UID.
func (r *itemRepo) DeleteItemForUser(ctx context.Context, id string, uid string) error {
	col := r.data.db.Collection("items")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return pb.InvalidItemIDError
	}

	_, err = col.DeleteOne(ctx, bson.M{"_id": objID, "uid": uid})

	return errors.Wrapf(err, "failed to delete %s", id)
}
