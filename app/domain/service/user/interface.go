package user

import (
	model "clean/app/domain/model/user"
	"context"
)

type UserService interface {
	RegisterData(ctx context.Context, Name, Password, Email string) *model.Register
}

type userService struct {

}
func NewUserService() UserService{
	return &userService{}
}
