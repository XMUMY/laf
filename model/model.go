package model

import (
	"time"

	"github.com/XMUMY/lost_found/proto/lost_found"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LostAndFoundBrief struct {
	Uid       string
	Type      lostfound.LostAndFoundType
	Name      string
	Timestamp time.Time
	Location  string
}

type LostAndFoundDetail struct {
	LostAndFoundBrief `bson:",inline"`
	Description       string
	Contacts          map[string]string
	Pictures          []string
}

type LostAndFoundItem struct {
	Id                 primitive.ObjectID `bson:"_id"`
	LostAndFoundDetail `bson:",inline"`
}
