package service

import (
	model "clean/app/domain/model/todo"
	"context"
)

type TodoService interface {
	CreateData(ctx context.Context, Matter, EndTime, FinishedCondition, Email string) (*model.Create, error)
	// GetAll(ctx context.Context, todo *model.HaveIDTodo) *model.HaveIDTodo
	DeleteData(ctx context.Context, Id,Email string) (*model.Delete)
	UpdateData(ctx context.Context, Email,Id,FinishedCondition,Note string) *model.Update
}
type todoService struct {
}

func NewTodoService() TodoService {
	return &todoService{}
}
