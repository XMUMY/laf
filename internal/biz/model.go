package biz

import (
	"time"

	v4 "github.com/XMUMY/lost_found/api/lost_found/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemBrief struct {
	Uid       string
	Type      v4.LostAndFoundType
	Name      string
	Timestamp time.Time
	Location  string
}

type ItemDetail struct {
	ItemBrief   `bson:",inline"`
	Description string
	Contacts    map[string]string
	Pictures    []string
}

type Item struct {
	Id         primitive.ObjectID `bson:"_id"`
	ItemDetail `bson:",inline"`
}
