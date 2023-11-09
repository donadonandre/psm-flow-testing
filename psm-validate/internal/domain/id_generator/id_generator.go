package id_generator

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"psm-validate/internal/infrastructure/database"
)

type IdGenerator struct {
	Collection string `json:"collection" bson:"collection"`
	CurrentId  uint32 `json:"current_id" bson:"current_id"`
}

func GetNextId(collectionName string) uint32 {
	client, collection := database.MongoClientGetCollection("idGenerator")
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	filter := bson.M{"collection": collectionName}

	var idGenerator IdGenerator
	err := collection.FindOne(context.Background(), filter).Decode(&idGenerator)
	if err != nil {
		idGenerator = IdGenerator{
			CurrentId:  1,
			Collection: collectionName,
		}
		_, err = collection.InsertOne(context.Background(), idGenerator)
		return 1
	}

	update := bson.M{"$set": bson.M{"collection": idGenerator.Collection, "current_id": idGenerator.CurrentId + 1}}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	return idGenerator.CurrentId + 1
}
