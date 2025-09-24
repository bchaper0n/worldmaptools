package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	//DBDriver          string
	DBSource          string
	MaxOpenConns      int
	MaxIdleConns      int
	ConnMaxIdleTime   time.Duration
	ConnectionTimeout time.Duration
}

type MongoClient struct {
	DB *mongo.Database
}

// creates a new database client with the given configuration.
func NewMongoClient(cfg Config) (*MongoClient, error) {
	// connect to db
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.DBSource))
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	// Ping the database to verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	db := client.Database("maps")

	return &MongoClient{DB: db}, nil
}

func (client *MongoClient) Close() error {
	if client.DB != nil {
		return client.DB.Client().Disconnect(context.TODO())
	}

	return nil
}
