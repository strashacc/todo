package mongodb

import (

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *Database

type Database struct {
	db *mongo.Database
}

func GetDatabase() *mongo.Database {
	return db.db
}

func StartConnection(uri string, name string) error {
	if db == nil {
		db = &Database{}
	}
	if uri == "" {
		return &NoURIError{}
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db.db = client.Database(name)
	return nil
}

type NoURIError struct {
	error
}
func (err *NoURIError) Error() string {
	return "Environment variable MONGO_URI was not set, connection is impossible"
}