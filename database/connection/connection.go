package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
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
func GetCollection(collection string) *mongo.Client {

	// Set up uri to connect to the database
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	// Connect to the database
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Remember to close connection
	defer client.Disconnect(ctx)

	// Check if the connection is stable
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// We can fetch all available databases from the connection
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	return client
}
