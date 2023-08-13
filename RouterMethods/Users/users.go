package userRouterMethods

import (
	userHelpers "github.com/Armaan-Singh-Dhillon/FurnitureStore/MongoHelpers/UserHelpers"
	"github.com/gorilla/mux"
)

func UserRouterMethods(router *mux.Router) {
	router.HandleFunc("/api/users",userHelpers.UserHelpers).Methods("GET")
}
