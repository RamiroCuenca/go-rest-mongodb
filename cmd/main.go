package main

import (
	"fmt"
	"net/http"

	"github.com/RamiroCuenca/go-rest-mongodb/database/connection"
)

func main() {
	// Get mongo client
	mongoClient := connection.GetMongoClient()

	// Set the database object
	db := mongoClient.Database("crud-database")

	// Set series and episodes collections
	seriesCollection := db.Collection("series")
	episodesCollection := db.Collection("episodes")

	// Get router
	mux := GetRouter()

	// Run locally
	http.ListenAndServe(":8000", mux)

	fmt.Println(db, seriesCollection, episodesCollection)
}
