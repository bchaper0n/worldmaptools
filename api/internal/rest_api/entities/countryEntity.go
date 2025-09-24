package entities

import "go.mongodb.org/mongo-driver/v2/bson"

type Country struct {
	ID           bson.ObjectID `bson:"_id"`
	Name         string        `json:"name,omitempty"`
	Abbreviation string        `json:"abbreviation,omitempty"`
	Capital      string        `json:"capital,omitempty"`
	Continent    string        `json:"continent,omitempty"`
	Flag         string        `json:"flag,omitempty"`
}
