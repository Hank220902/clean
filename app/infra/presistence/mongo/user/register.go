package user

import (
	model "clean/app/domain/model/user"
	"clean/app/infra/presistence/mongo"
	"context"
	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	success     int = 1
	fail        int = 2
	tokenError  int = 3
	emailExists int = 4
)

func (r *repo) Register(ctx context.Context, data *model.Register) int {

	if haveEmail(ctx,data.Email)==emailExists{
		return emailExists
	}

	insertOne, err := mongo.UserCollection.InsertOne(ctx, data)
	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
	if err != nil {
		log.Fatal(err)
	}
	return success
}

func haveEmail(ctx context.Context, email string) int {
	var reqLogin model.Register
	fmt.Println(reqLogin.Password)

	var resultLogin model.Register
	filter := bson.D{{Key: "email", Value: email}}
	err := mongo.UserCollection.FindOne(ctx, filter).Decode(&resultLogin)
	if err != nil {
		return success
	} else {
		return emailExists
	}
}
