package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LykedUploads struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserID      string             `bson:"user_id" json:"user_id"` // Store User UUID as a string
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	VideoLink   string             `bson:"video_link" json:"video_link"`
	Folders     []string           `bson:"folders" json:"folders"`
	Tags        []string           `bson:"tags" json:"tags"`
}
