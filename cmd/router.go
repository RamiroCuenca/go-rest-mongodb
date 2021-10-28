package main

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	mux := mux.NewRouter()

	return mux
}
