package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"encoding/json"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/joho/godotenv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Name         string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Capital      string `json:"capital,omitempty"`
	Continent    string `json:"continent,omitempty"`
}

// https://go.dev/doc/tutorial/web-service-gin
func getCountries(c *gin.Context) {
	println("Getting countries...")
	client := connect()
	coll := client.Database("map").Collection("countries")

	filter := bson.D{}
	sort := bson.D{{"name", 1}}
	opts := options.Find().SetSort(sort)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}

	results := []Country{}

	for cursor.Next(context.TODO()) {
		var result Country
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%+v\n", result)
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, results)
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
		panic(err)
	}

	database := client.Database("map")

	err = database.CreateCollection(context.TODO(), "countries")
	if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}

	return client
}

func initCountries(client *mongo.Client) {
	database := client.Database("map")
	coll := database.Collection("countries")

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	err := database.Collection("countries").Drop(context.TODO())
	if err != nil {
		panic(err)
	}

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

	// for _, id := range result.InsertedIDs {
	// 	fmt.Printf("Inserted document with _id: %v\n", id)
	// }
}

func main() {

	reset := true

	if reset {
		client := connect()
		initCountries(client)
		println("Inserted countries")
	}

	router := gin.Default()
	router.GET("/countries", getCountries)

	router.Run("localhost:8080")

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
}
