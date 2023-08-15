package productRouterMethods

import (
	producthelpers "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoHelpers/ProductHelpers"
	"github.com/gorilla/mux"
)

func ProductRouterMethods(router *mux.Router) {
	router.HandleFunc("/api/products/getall", producthelpers.GetProducts).Methods("GET").Queries("sortPrice","{sortPrice}")
	router.HandleFunc("/api/products/create", producthelpers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/get/{id}", producthelpers.GetProductById).Methods("GET")
}
