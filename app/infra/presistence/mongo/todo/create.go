package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"

	"log"
)
const (
	success     int = 1
	fail        int = 2
	tokenError  int = 3
	emailExists int = 4
)

func (r *repo) Create(ctx context.Context, data *model.Create) error {

	_, err := mongo.TodoCollection.InsertOne(ctx, data)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
