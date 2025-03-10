package routes

import (
	"gorm/controllers"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")
}
