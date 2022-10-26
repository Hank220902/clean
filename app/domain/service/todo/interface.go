package service

import (
	model "clean/app/domain/model/todo"
	"context"
)

type TodoService interface {
	NewTodo(ctx context.Context, Matter, EndTime, FinishedCondition, Email string) (*model.Todo, error)
	// GetTodo(ctx context.Context, todo *model.HaveIDTodo) *model.HaveIDTodo
	// DeleteTodo(ctx context.Context, todo *model.HaveIDTodo) string
	// UpdateTodo(ctx context.Context, todo *model.HaveIDTodo) string
}
type todoService struct {
}

func NewTodoService() TodoService {
	return &todoService{}
}
