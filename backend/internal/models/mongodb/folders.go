package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Folder struct {
	ID      bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"user_id,omitempty" json:"user_id"`
	Name    string             `bson:"name,omitempty" json:"name"`
	PostIDs []string           `bson:"post_ids,omitempty" json:"post_ids"`
}
