package producthelpers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	dbConnection "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoConnector"
	"github.com/Armaan-Singh-Dhillon/FurnitureStore/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = dbConnection.ProductCollection

// Helpers
func getAllProducts(name string, category string, sortPrice int) []primitive.M {
	filter := bson.M{}

	options := options.Find().SetSort(bson.M{"price": sortPrice})

	if name != "" {
		filter["name"] = name
	}

	if category != "" {
		filter["category"] = category
	}

	mongoCursor, err := collection.Find(context.Background(), filter, options)
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

func insertProduct(product models.Product) {
	_, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}
}

func getProductById(stringId string) (primitive.M, error) {
	id, err := primitive.ObjectIDFromHex(stringId)
	if err != nil {
		return nil, err
	}
	var result bson.M

	collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)

	return result, nil

}

// Controller
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")
	sortPrice := r.URL.Query().Get("sortPrice")
	sortPriceInt, err := strconv.Atoi(sortPrice)
	if err != nil {
		log.Fatal(err)
	}
	products := getAllProducts(name, category, sortPriceInt)

	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	insertProduct(product)
	json.NewEncoder(w).Encode(product)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	singleProduct, err := getProductById(params["id"])
	if err != nil {

		json.NewEncoder(w).Encode("provided id is not in correct format")
		return
	}

	json.NewEncoder(w).Encode(singleProduct)

}
