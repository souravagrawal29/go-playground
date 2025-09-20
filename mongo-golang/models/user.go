package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


type User struct {
	Id		bson.ObjectID	`json:"id" bson:"_id"`
	Name 	string			`json:"name" bson:"name"`
	Gender 	string			`json:"gender" bson:"gender"`
	Age		int				`json:"age" bson:"age"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}


var client *mongo.Client

func init() {
	dataSource := "mongodb://127.0.0.1:27017/"
	c, err := mongo.Connect(options.Client().ApplyURI(dataSource))
	if err != nil {
		panic(err)
	}
	client = c
}

func GetClient() *mongo.Client {
	return client
}