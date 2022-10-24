package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"log"
)

func Create(ctx context.Context, data *model.Todo) error {
	_, err := mongo.TodoCollection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
