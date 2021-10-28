package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Series struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Category string             `bson:"category,omitempty"`
	Cast     []string           `bson:"cast,omitempty"`
}
