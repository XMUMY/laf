package dao

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type Dao struct {
	mgo *mongo.Client
	db  *mongo.Database
}

func New() *Dao {
	client := InitMongo()

	return &Dao{
		mgo: client,
		db:  client.Database(os.Getenv("DB_NAME")),
	}
}
