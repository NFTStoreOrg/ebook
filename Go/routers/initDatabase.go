package routers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var database *mongo.Client

// Initial database
func InitDB() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//	Connect to mongodb
	database, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//	Check connection
	err = database.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *mongo.Client {
	if database == nil {
		InitDB()
	}
	return database
}
