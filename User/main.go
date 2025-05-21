package main

import (
	"log"
	"net/http"
	"user/handlers"
	"user/storage"

	"github.com/gorilla/mux"
)

func main() {
	storage.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
