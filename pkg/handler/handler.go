package handler

import (
	. "chat-example/app"
	"chat-example/pkg/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
)

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}

func GetMessages(w http.ResponseWriter, r *http.Request) {

	//wsConn, err := upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//})

	//for {
	//	// Read message from browser
	//	msgType, msg, err := wsConn.ReadMessage()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	// Print the message to the console
	//	fmt.Printf("%s sent: %s\n", wsConn.RemoteAddr(), string(msg))
	//	msgClean := strings.TrimSpace(string(msg))
	//	if len(msgClean) > 0 {
	//
	//		db, dbClient, ctx := getDbConnect()
	//		defer dbClient.Disconnect(ctx)
	//		coll := db.Collection("messages")
	//		message := Message{
	//			Text:      string(msg),
	//			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	//		}
	//		insertResult, err := coll.InsertOne(ctx, message)
	//
	//		if err != nil {
	//			panic(err)
	//		}
	//		fmt.Println(insertResult.InsertedID)
	//
	//		// Write message back to browser
	//		if err = wsConn.WriteMessage(msgType, msg); err != nil {
	//			log.Println(err)
	//			return
	//		}
	//	} else {
	//		log.Println("Send empty string!")
	//	}
	//}
}
func GetMainPage(w http.ResponseWriter, r *http.Request) {
	db, dbClient, ctx := repository.GetDbConnect()
	defer dbClient.Disconnect(ctx)
	msgCollection := db.Collection("messages")
	var messages []Message
	cursor, err := msgCollection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &messages); err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	tmpl := template.Must(template.ParseFiles("websockets.html"))

	type PageData struct {
		PageTitle string
		Messages  []Message
	}

	tmpl.Execute(w, PageData{Messages: messages})
	//http.ServeFile(w, r, "websockets.html")
}
