package userHelpers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	dbConnection "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoConnector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = dbConnection.UserCollection

func UserHelpers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	mongoCursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var users []primitive.M

	for mongoCursor.Next(context.Background()){
		var user bson.M
		err := mongoCursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)

	}
	json.NewEncoder(w).Encode(users)
}
