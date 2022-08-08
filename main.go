package main

import (
	"crud-go/controllers"
	"crud-go/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.GetAllProducts).Methods("GET")
	// router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}

func main() {
	database.Connect("user:1234@tcp(127.0.0.1:3306)/test_db")
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	RegisterProductRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", "8080"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", "8080"), router))
}
