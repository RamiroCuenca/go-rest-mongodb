package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Series struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Category string             `json:"category,omitempty" bson:"category,omitempty"`
	Cast     []string           `json:"cast,omitempty" bson:"cast,omitempty"`
}
