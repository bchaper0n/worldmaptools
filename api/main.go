package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"encoding/json"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/joho/godotenv"
)

type Country struct {
	Name         string `json:"country,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Capital      string `json:"capital city,omitempty"`
	Continent    string `json:"continent,omitempty"`
}

func connect() *mongo.Client {
	// get mongo auth from env vars
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := "mongodb://" + os.Getenv("MONGO_ROOT_USERNAME") + ":" + os.Getenv("MONGO_ROOT_PASSWORD") + "@localhost:27017/"

	// connect to db
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		println("hi1")
		panic(err)
	}

	database := client.Database("map")

	err = database.CreateCollection(context.TODO(), "countries")
	if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}

	err = database.Collection("countries").Drop(context.TODO())
	if err != nil {
		panic(err)
	}

	return client
}

func main() {

	client := connect()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			println("hi2")
			panic(err)
		}
	}()

	database := client.Database("map")
	coll := database.Collection("countries")

	countriesJSON, err := os.ReadFile("./data/countries.json")
	if err != nil {
		panic(err)
	}

	var countries []Country

	err = json.Unmarshal([]byte(countriesJSON), &countries)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// for i, c := range countries {
	// 	fmt.Printf("Country %d and %s: %s with %s\n", i, c.CountryId, c.Name, c.Capital)
	// }

	result, err := coll.InsertMany(context.TODO(), countries)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))

	for _, id := range result.InsertedIDs {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}
