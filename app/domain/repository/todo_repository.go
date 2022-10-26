package repository

import (
	model "clean/app/domain/model/todo"
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) error
	// GetAll(ctx context.Context, todo *model.HaveIDTodo) (*model.HaveIDTodo,error)
}
