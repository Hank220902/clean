package repository

import ("context"
	model "clean/app/domain/model/todo"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo)error
}
