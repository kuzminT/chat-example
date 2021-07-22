package pkg

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Text string	`bson:"text,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}
