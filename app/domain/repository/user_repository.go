package repository

import (
	model "clean/app/domain/model/user"
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, todo *model.Register) int
}