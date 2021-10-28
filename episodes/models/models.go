package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Episode struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Serie       primitive.ObjectID `json:"serie,omitempty" bson:"serie,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Duration    int32              `json:"duration,omitempty" bson:"duration,omitempty"`
}
