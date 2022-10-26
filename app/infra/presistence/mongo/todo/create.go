package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"fmt"
	"log"
)

func (r *repo) Create(ctx context.Context, data *model.Todo) error {
	fmt.Println(r)
	_, err := mongo.TodoCollection.InsertOne(ctx, data)
	fmt.Println("qweqeqe")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
