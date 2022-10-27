package user

import (
	model "clean/app/domain/model/user"
	"context"
)

func (u *userService) RegisterData(ctx context.Context, Name, Password, Email string) *model.Register{
	var RegisterData = new(model.Register)
	RegisterData.Email = Email
	RegisterData.Name = Name
	RegisterData.Password= Password
	return RegisterData
}
