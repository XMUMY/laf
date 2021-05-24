package data

import (
	"context"
	"os"
	"time"

	"github.com/XMUMY/lost_found/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewItemRepo)

// Data .
type Data struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	mongoOptions := options.Client().ApplyURI(os.Getenv("DB_ADDR"))
	mongoOptions.Auth = &options.Credential{
		AuthSource: os.Getenv("DB_NAME"),
		Username:   os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
	}

	// Connect to mongodb.
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		logger.Log("msg", "closing the data resources")
		client.Disconnect(context.TODO())
	}
	return &Data{
		client: client,
		db:     client.Database(os.Getenv("DB_NAME")),
	}, cleanup, nil
}
