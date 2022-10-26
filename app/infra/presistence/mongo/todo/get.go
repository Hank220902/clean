package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetAll(ctx context.Context, email string) ([]*model.HaveIDTodo, error) {
	var results []*model.HaveIDTodo
	filter := bson.D{{Key: "email", Value: email}}
	
	cur, err := mongo.TodoCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.All(ctx, &results); err != nil {

	}

	return results, nil

}
