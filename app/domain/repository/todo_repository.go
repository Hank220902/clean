package repository

import (
	model "clean/app/domain/model/todo"
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Create) error
	GetAll(ctx context.Context, email string) ([]*model.Get, error)
	Delete(ctx context.Context, delete *model.Delete) int
	Update(ctx context.Context, todo *model.Update) int
}
