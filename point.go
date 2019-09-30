package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Point is a simple type with a location for geospatial
// queries.
type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name"`
	Location Location           `json:"location"`
}
