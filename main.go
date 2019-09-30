package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if err := createDBSession(); err != nil {
		fmt.Println(err)
		return
	}

	if err := createIndex(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected")

	p := Point{
		Name:     "Central Park",
		Location: NewPoint(-73.97, 40.77),
	}

	err := AddPoint(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	points, err := GetPointsByDistance(NewPoint(-73.97, 40.77), 50)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(points)
}

// AddPoint adds a new point to the collection.
func AddPoint(point Point) error {
	coll := conn.Database(DBName).Collection(PointCollection)

	point.ID = primitive.NewObjectID()

	insertResult, err := coll.InsertOne(ctx, point)
	if err != nil {
		fmt.Printf("Could not insert new Point. Id: %s\n", point.ID)
		return err
	}

	fmt.Printf("Inserted new Point. ID: %s\n", insertResult.InsertedID)
	return nil
}

// GetPointsByDistance gets all the points that are within the maximum
// distance provided in meters.
func GetPointsByDistance(location Location, distance int) ([]Point, error) {
	coll := conn.Database(DBName).Collection(PointCollection)
	var results []Point

	filter := bson.D{
		{"location",
			bson.D{
				{"$near", bson.D{
					{"$geometry", location},
					{"$maxDistance", distance},
				}},
			}},
	}

	cur, err := coll.Find(ctx, filter)

	if err != nil {
		return []Point{}, err
	}

	for cur.Next(ctx) {
		var elem Point
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Could not decode Point")
			return []Point{}, err
		}

		results = append(results, elem)
	}

	return results, nil
}
