package router

import (
	citiesRouterMethods "github.com/Armaan-Singh-Dhillon/FurnitureStore/RouterMethods/Cities"
	productRouterMethods "github.com/Armaan-Singh-Dhillon/FurnitureStore/RouterMethods/Products"
	userRouterMethods "github.com/Armaan-Singh-Dhillon/FurnitureStore/RouterMethods/Users"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	productRouterMethods.ProductRouterMethods(router)
	userRouterMethods.UserRouterMethods(router)
	citiesRouterMethods.CitiesRouterMethods(router)

	return router
}
