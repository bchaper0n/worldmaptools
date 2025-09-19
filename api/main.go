package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/joho/godotenv"
)

func main() {

	// get mongo auth from env vars
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := "mongodb://" + os.Getenv("MONGO_ROOT_USERNAME") + ":" + os.Getenv("MONGO_ROOT_PASSWORD") + "@localhost:27017/"

	// connect to db
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database := client.Database("countries")

	err = database.CreateCollection(context.TODO(), "example_collection")
	if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}
}
