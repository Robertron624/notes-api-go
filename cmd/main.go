package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-api/db"
	"notes-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world!!!")

	DB := db.Init()

	h := handlers.New(DB)

	router := mux.NewRouter()

	//Health check
	router.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)

	log.Println("Server listening on port 8080")

	http.ListenAndServe(":8080", router)

}
