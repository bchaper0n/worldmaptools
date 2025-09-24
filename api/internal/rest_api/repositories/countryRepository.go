package repositories

import (
	"worldmaptools/wmt/internal/rest_api/database"
	"worldmaptools/wmt/internal/rest_api/entities"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var coll = "countries"

type Country struct {
	database.BaseMongoRepository[entities.Country]
}

func NewCountryRepository(db *mongo.Database) *Country {
	return &Country{
		BaseMongoRepository: database.BaseMongoRepository[entities.Country]{DB: db},
	}
}

func mapCountry(row *mongo.Cursor, c *entities.Country) error {
	return row.Decode(&c)
}

func (r *Country) GetAllCountries() ([]*entities.Country, error) {

	params := &database.BaseMongoParams{
		Collection: coll,
		Filter:     bson.D{},
		Options:    options.Find().SetSort(bson.D{{"name", 1}}),
	}

	return r.SelectMultiple(mapCountry, params)
}
