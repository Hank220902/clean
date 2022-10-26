package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) Delete(ctx context.Context, data *model.Delete) int {

	id, _ := primitive.ObjectIDFromHex(data.ID)

	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: data.Email}}

	deleteResult, err := mongo.TodoCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	if deleteResult.DeletedCount == 0 {
		return 2
	}

	return 1
}
