package handlers

import (
	"encoding/json"
	"net/http"
	"user/models"
	"user/storage"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User

	if err := storage.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	storage.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	storage.DB.Delete(&models.User{}, id)
	w.WriteHeader(http.StatusNoContent)
}
