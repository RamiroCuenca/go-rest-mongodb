package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RamiroCuenca/go-rest-mongodb/common"
	"github.com/RamiroCuenca/go-rest-mongodb/database/connection"
	"github.com/RamiroCuenca/go-rest-mongodb/series/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
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

	// 4. InsertOne document - result will contain InsertedID int
	seriesCollection.InsertOne(ctx, serie)

	// result, err := seriesCollection.InsertOne(ctx, serie)

	// serie.ID = result.InsertedID // Does not work :(

	// 5. Encode the result (Contains the id)
	// json.NewEncoder(w).Encode(result) // Return the ObjectID
	json, _ := json.Marshal(serie) // Return the serie but without ID

	// 6. Send response
	common.SendResponse(w, http.StatusOK, []byte(json))
}

// Fetch all series
func GetAll(w http.ResponseWriter, r *http.Request) {
	// 1. Connect to the database and setup collection variable
	// 2. Create a Series Array where we can store data
	// 3. Setup context
	// 4. Execute the query (Remember to close cursor)
	// 5. Insert data from cursor on array
	// 6. Send response

	// 1.
	db := connection.GetMongoClient()
	collection := db.Database("crud-database").Collection("series")

	// 2.
	var seriesArr []models.Series

	// 3.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// 4.
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		common.SendError(w, http.StatusInternalServerError, []byte(`{"message": `+err.Error()+`}`))
		return
	}
	defer cursor.Close(ctx)

	// 5.
	for cursor.Next(ctx) {
		var serie models.Series

		cursor.Decode(&serie)

		seriesArr = append(seriesArr, serie)
	}

	if err = cursor.Err(); err != nil {
		common.SendError(w, http.StatusInternalServerError, []byte(`{"message": `+err.Error()+`}`))
		return
	}

	// 6.
	// json.NewEncoder(w).Encode(seriesArr)
	json, err := json.Marshal(seriesArr)

	resp := fmt.Sprintf(`{
		"result": %s
	}`, json)

	common.SendResponse(w, http.StatusOK, []byte(resp))
}

// Fetch one serie by id
func GetById(w http.ResponseWriter, r *http.Request) {
	// 1. Obtain the id from the url params
	// 2. Connect to the database and collection
	// 3. Create a serie object where store data
	// 4. Setup context
	// 5. Execute query
	// 6. Store data on serie object
	// 7. Send response

	// 1.
	urlParam := r.URL.Query().Get("id") // Return a string, should convert it in primitive.ID
	id, err := primitive.ObjectIDFromHex(urlParam)
	if err != nil {
		common.SendError(w, http.StatusBadRequest, []byte(`{"message": `+err.Error()+`}`))
		return
	}

	// 2.
	db := connection.GetMongoClient()
	collection := db.Database("crud-database").Collection("series")

	// 3.
	var serie models.Series

	// 4.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// 5.
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&serie)

	// json.NewEncoder(w).Encode(serie) // Ugly

	json, _ := json.Marshal(serie)

	data := fmt.Sprintf(`{
			"result": %s
		}`, json)

	// 6.
	common.SendResponse(w, http.StatusOK, []byte(data))
}

// Delete by id
func Delete(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch the id from url
	// 2. Connect to the mongo database and select collection
	// 3. Setup context
	// 4. Execute query
	// 5. Send response

	// 1.
	urlParam := r.URL.Query().Get("id")

	// It's a string, should turn it into a primitive.ID
	id, err := primitive.ObjectIDFromHex(urlParam)
	if err != nil {
		common.SendError(w, http.StatusBadRequest, []byte(`{"message": `+err.Error()+`}`))
		return
	}

	// 2.
	client := connection.GetMongoClient()
	collection := client.Database("crud-database").Collection("series")

	// 3.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// 4.
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		common.SendError(w, http.StatusBadRequest, []byte(`{"message": `+err.Error()+`}`))
		return
	}

	if result.DeletedCount == 0 {
		data := []byte(`{
			"message": "Query executed succesfully but could not find any document that matchs with received _id",
		}`)
		common.SendResponse(w, http.StatusOK, data)
		return
	} else {
		data := []byte(`{
			"message": "Document deleted successfully",
		}`)
		common.SendResponse(w, http.StatusOK, data)
		return
	}

}
