package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// DBName Database name.
const DBName = "geolocation"

var (
	conn       *mongo.Client
	ctx               = context.Background()
	connString string = "mongodb://127.0.0.1:27017/glottery"
)

// Database collections.
var (
	PointCollection = "points"
)

// createDBSession Create a new connection with the database.
func createDBSession() error {
	var err error

	conn, err = mongo.Connect(ctx, options.Client().
		ApplyURI(connString))
	if err != nil {
		return err
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func createIndex() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := conn.Database(DBName)
	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)

	// Index to location 2dsphere type.
	pointIndexModel := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys:    bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}

	pointIndexes := db.Collection(PointCollection).Indexes()

	_, err := pointIndexes.CreateOne(ctx, pointIndexModel, indexOpts)
	if err != nil {
		return err
	}

	return nil
}
