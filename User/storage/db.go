package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB variables
var Client *mongo.Client
var UserCollection *mongo.Collection

func InitDB() {
	// MongoDB connection URI
	uri := "mongodb://localhost:27017"

	// MongoDB client options
	clientOpts := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Check the connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	// Get user collection from database
	UserCollection = Client.Database("todo_db").Collection("users")
	log.Println("MongoDB initialized and 'users' collection is ready.")
}
