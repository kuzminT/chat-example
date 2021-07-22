package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const ConnectionUrl string = "mongodb://admin:test@localhost:27017"

func GetDbConnect() (*mongo.Database, *mongo.Client, context.Context) {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(ConnectionUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
	//	log.Fatalf("Connection with db is absent: %s", err)
	//	return nil
	//}
	return dbClient.Database("chat"), dbClient, ctx
}
