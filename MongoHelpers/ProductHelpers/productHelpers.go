package producthelpers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dbConnection "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoConnector"
	errorPackage "github.com/Armaan-Singh-Dhillon/FurnitureStore/errors"
	"github.com/Armaan-Singh-Dhillon/FurnitureStore/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = dbConnection.ProductCollection

// Helpers
func getAllProducts(name string, category string, sortPrice int) (_ []primitive.M, findError error, decodeError error) {
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
		return nil, err, nil
	}
	var products []primitive.M
	for mongoCursor.Next(context.Background()) {
		var product bson.M
		err := mongoCursor.Decode(&product)
		if err != nil {
			return nil, nil, err
		}
		products = append(products, product)
	}
	defer mongoCursor.Close(context.Background())
	fmt.Println(products)
	return products, nil, nil
}

func insertProduct(product models.Product) error {
	_, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		return err
	}
	return nil
}

func getProductById(stringId string) (_ primitive.M, err error) {
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorPackage.APIError{
			Code:    http.StatusBadRequest,
			Message: "Query Parameter sortPrice expects 1 or -1 as a value",
		})
		return
	}
	products, findError, decodeError := getAllProducts(name, category, sortPriceInt)
	if findError != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorPackage.APIError{
			Code:    http.StatusInternalServerError,
			Message: "internal server error - cannot fetch data from the servers",
		})
		return

	}
	if decodeError != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorPackage.APIError{
			Code:    http.StatusInternalServerError,
			Message: "cannot decode the datatype from the database - try again later",
		})
		return
	}

	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorPackage.APIError{
			Code:    http.StatusBadRequest,
			Message: "cannot decode the datatype from the database - try again later",
		})
		return
	}
	err = insertProduct(product)
	if err := insertProduct(product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorPackage.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error: Failed to insert data into the database",
		})
		return
	}
	json.NewEncoder(w).Encode(product)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	result, err := getProductById(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorPackage.ErrBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)

}
