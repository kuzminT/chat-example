package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const ConnectionUrl string = "mongodb://admin:test@localhost:27017"

func getDbConnect () (*mongo.Database, *mongo.Client, context.Context) {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(ConnectionUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//databases, err := dbClient.ListDatabaseNames(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err, 1)
	//}
	//fmt.Println(databases)
	//if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
	//	log.Fatalf("Connection with db is absent: %s", err)
	//	return nil
	//}
	return dbClient.Database("chat"), dbClient, ctx
}


func main() {

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {

		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
			return
		}

		for {
			// Read message from browser
			msgType, msg, err := wsConn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", wsConn.RemoteAddr(), string(msg))

			db, dbClient, ctx := getDbConnect()
			defer dbClient.Disconnect(ctx)
			coll := db.Collection("messages")
			message :=  Message{
				Text: string(msg),
				CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
			}
			insertResult, err := coll.InsertOne(ctx, message)

			if err != nil {
				panic(err)
			}
			fmt.Println(insertResult.InsertedID)

			// Write message back to browser
			if err = wsConn.WriteMessage(msgType, msg); err != nil {
				log.Println(err)
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db, dbClient, ctx := getDbConnect()
		defer dbClient.Disconnect(ctx)
		msgCollection := db.Collection("messages")
		var messages []Message
		cursor, err := msgCollection.Find(ctx, bson.M{})
		if err != nil {
			panic(err)
		}
		if err = cursor.All(ctx, &messages); err != nil {
			panic(err)
		}
		fmt.Printf("Messages: %s", messages)
		http.ServeFile(w, r, "websockets.html")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
