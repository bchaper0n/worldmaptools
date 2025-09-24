package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BaseMongoRepository[T any] struct {
	DB *mongo.Database
}

type BaseMongoParams struct {
	Collection string
	Filter     bson.D
	Options    *options.FindOptionsBuilder
}

func (repo *BaseMongoRepository[T]) SelectMultiple(mapRow func(*mongo.Cursor, *T) error, params *BaseMongoParams) ([]*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := repo.DB.Collection(params.Collection).Find(context.TODO(), params.Filter, params.Options)
	if err != nil {
		panic(err)
	}

	defer func(cursor *mongo.Cursor) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor)

	var list []*T

	// Loop through rows, using Scan to assign column data to struct fields.
	for cursor.Next(ctx) {
		var t T
		if err := mapRow(cursor, &t); err != nil {
			return nil, err
		}
		list = append(list, &t)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
