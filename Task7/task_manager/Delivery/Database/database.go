package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct{
}

type DB interface{
	Connection(URI string) *mongo.Client 
	ConnectToDatabase()*mongo.Client
	CreateDb(collectionName string)*mongo.Collection
}

func NewDatabase()*Db{
	return &Db{}
}

func (db *Db)Connection(URI string) *mongo.Client {
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

func (db *Db)ConnectToDatabase()*mongo.Client{
	connection := db.Connection("mongodb://localhost:27017")
	return connection
} 

func (db *Db)CreateDb(collectionName string)*mongo.Collection{
	connection := db.ConnectToDatabase()
	collection := connection.Database("task_manager").Collection(collectionName)
	return collection
} 
