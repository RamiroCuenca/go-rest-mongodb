package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Episode struct {
	ID          primitive.ObjectID `bson:"id,omitempty"`
	Serie       primitive.ObjectID `bson:"serie,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}
