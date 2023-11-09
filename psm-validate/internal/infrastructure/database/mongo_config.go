package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func mongoConnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:mongodb@localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("Error while connect to mongo", err)
	}

	return client
}

func MongoClientGetCollection(collectionName string) (*mongo.Client, *mongo.Collection) {
	client := mongoConnect()
	return client, client.Database("psm-validate").Collection(collectionName)
}
