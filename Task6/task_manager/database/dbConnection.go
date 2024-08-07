package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection(URI string) *mongo.Client {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()
	connection, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	err = connection.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer connection.Disconnect(context.TODO())
	return connection
}

func ConnectToDatabase()*mongo.Client{
	connection := Connection("mongodb://localhost:27017")
	return connection
}