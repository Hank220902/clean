package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"time"
)

var TodoCollection *mongo.Collection
var UserCollection *mongo.Collection


func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var ctx = context.TODO()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	//defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)



	TodoCollection = client.Database("todolist").Collection("todolists")
	UserCollection = client.Database("todolist").Collection("users")
	fmt.Println("Collection instance created!")

}
