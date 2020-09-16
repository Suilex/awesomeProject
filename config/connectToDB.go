package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var collect *mongo.Collection

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://Admin:AdminPassword@cluster0.w3oor.gcp.mongodb.net/dbTest?retryWrites=true&w=majority",
	))

	collect = client.Database("dbTest").Collection("collectionT")
}

var GetDB = func () *mongo.Collection {
	return collect
}
