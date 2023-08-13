package producthelpers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	dbConnection "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoConnector"
	"github.com/Armaan-Singh-Dhillon/FurnitureStore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = dbConnection.ProductCollection

//Helpers
func getAllProducts() []primitive.M {
	mongoCursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var products []primitive.M
	for mongoCursor.Next(context.Background()) {
		var product bson.M
		err := mongoCursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	defer mongoCursor.Close(context.Background())
	return products
}

func insertProduct(product models.Product)  {
	_, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}
}
// Controller
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	products := getAllProducts()

	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var product models.Product
	err:=json.NewDecoder(r.Body).Decode(&product)
	if(err!=nil){
		log.Fatal(err)
	}
	insertProduct(product)
	json.NewEncoder(w).Encode(product)


}
