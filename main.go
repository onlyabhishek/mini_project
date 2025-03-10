package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm/database"
	"gorm/models"
	"gorm/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDatabase()
	models.MigrateUserTable(database.DB)

	router := mux.NewRouter()
	routes.SetupUserRoutes(router)

	port := ":8081"
	fmt.Println("Server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
