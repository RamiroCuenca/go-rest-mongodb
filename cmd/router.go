package main

import (
	series "github.com/RamiroCuenca/go-rest-mongodb/series/controllers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	// Set up the router
	mux := mux.NewRouter()

	// Handlers
	mux.HandleFunc("/series", series.Create).Methods("POST")

	return mux
}
