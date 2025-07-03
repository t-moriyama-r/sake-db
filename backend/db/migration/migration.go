package main

import (
	"backend/db/repository/liquorRepository"
	"backend/util/helper"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// go run db/migration/migration.go

func main() {
	helper.LoadEnv()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	dbName := os.Getenv("MAIN_DB_NAME")
	collection := client.Database(dbName).Collection(liquorRepository.CollectionName)
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			log.Printf("Failed to decode document: %v\n", err)
			continue
		}

		_id := doc["_id"].(primitive.ObjectID)

		setFields := bson.M{}
		for _, field := range []string{"rate1_users", "rate2_users", "rate3_users", "rate4_users", "rate5_users"} {
			if doc[field] == nil {
				setFields[field] = []primitive.ObjectID{}
			}
		}

		if len(setFields) > 0 {
			update := bson.M{"$set": setFields}
			_, err := collection.UpdateByID(ctx, _id, update)
			if err != nil {
				log.Printf("Failed to update document %v: %v\n", _id.Hex(), err)
			} else {
				fmt.Printf("Updated document %v with fields: %v\n", _id.Hex(), setFields)
			}
		}
	}
}
