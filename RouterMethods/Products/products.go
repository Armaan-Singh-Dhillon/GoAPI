package productRouterMethods

import (
	producthelpers "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoHelpers/ProductHelpers"
	"github.com/gorilla/mux"
)

func ProductRouterMethods(router *mux.Router) {
	router.HandleFunc("/api/products", producthelpers.GetProducts).Methods("GET")
}
