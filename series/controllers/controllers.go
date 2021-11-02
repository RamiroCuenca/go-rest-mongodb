package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/RamiroCuenca/go-rest-mongodb/common"
	"github.com/RamiroCuenca/go-rest-mongodb/database/connection"
	"github.com/RamiroCuenca/go-rest-mongodb/series/models"
)

// Create a new serie
func Create(w http.ResponseWriter, r *http.Request) {
	// 1. Create a Serie object where we can store sent values
	var serie models.Series

	// 2. Decode the sent data in serie
	err := json.NewDecoder(r.Body).Decode(&serie)
	if err != nil {
		data := []byte("Try sending valid values")
		common.SendError(w, http.StatusBadRequest, data)
		return
	}

	// 3. Get the mongo client and set up collection var
	mongoClient := connection.GetMongoClient()
	seriesCollection := mongoClient.Database("crud-database").Collection("series")

	// Set context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// 4. InsertOne document
	result, err := seriesCollection.InsertOne(ctx, serie)

	// 5. Encode the result (Contains the id)
	json.NewEncoder(w).Encode(result)
	// json, _ := json.Marshal(result.InsertedID)

	// 6. Send response
	// common.SendResponse(w, http.StatusOK, []byte(json))

}
