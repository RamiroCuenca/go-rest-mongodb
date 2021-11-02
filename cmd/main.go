package main

import (
	"net/http"
)

func main() {
	// Get mongo client
	// MongoClient := connection.GetMongoClient()

	// // Set the database object
	// DB := MongoClient.Database("crud-database")

	// // Set series and episodes collections
	// seriesCollection := db.Collection("series")
	// episodesCollection := db.Collection("episodes")

	// Get router
	mux := GetRouter()

	// Run locally
	http.ListenAndServe(":8000", mux)
}
