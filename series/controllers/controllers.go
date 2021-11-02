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
