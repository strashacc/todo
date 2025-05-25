package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/strashacc/todo/Notification/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserNotificationsREST(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDstr := r.URL.Query().Get("userId")
		userID, err := primitive.ObjectIDFromHex(userIDstr)
		if err != nil {
			http.Error(w, "Invalid userId", 400)
			return
		}
		filter := bson.M{"userId": userID}
		cursor, err := db.Collection("notifications").Find(context.Background(), filter)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var notifs []models.Notification
		cursor.All(context.Background(), &notifs)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notifs)
	}
}
