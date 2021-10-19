package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Env variables
const (
	user     = "admin"
	password = "password"
	host     = "localhost"
	port     = "27017"
	database = "crud-database"
)

// Get a collection from the database
func GetCollection(collection string) *mongo.Collection {

	// Set up uri to connect to the database
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	// Connect to the database
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return client.Database(database).Collection(collection)
}
