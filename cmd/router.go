package main

import (
	series "github.com/RamiroCuenca/go-rest-mongodb/series/controllers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	// Set up the router
	mux := mux.NewRouter()

	// Handlers
	mux.HandleFunc("/series/create", series.Create).Methods("POST")
	mux.HandleFunc("/series/getall", series.GetAll).Methods("GET")
	mux.HandleFunc("/series/getbyid", series.GetById).Methods("GET")

	return mux
}
