package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Update(ctx context.Context, data *model.Update) int {
	id, _ := primitive.ObjectIDFromHex(data.ID)
	filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: data.Email}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "finishedCondition", Value: data.FinishedCondition},
			{Key: "note", Value: data.Note},
			{Key: "updateTime", Value: time.Now()},
		},
		}}
	result, err := mongo.TodoCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return fail
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
		return success
	}
	return fail

}
