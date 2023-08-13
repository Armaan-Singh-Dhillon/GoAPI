package router

import (
	dbConnection "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoConnector"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/getall",dbConnection.GetProducts).Methods("GET")

	return router
}
