package login

import (
	repository "clean/app/domain/repository"
	service "clean/app/domain/service/user"
	"context"
)

type loginUsecase struct {
	userService service.UserService
	userRepo    repository.UserRepository
}

func NewLgoinUsecase(userRepo repository.UserRepository, userService service.UserService) LoginUsecase {
	return &loginUsecase{
		userRepo:    userRepo,
		userService: userService,
	}
}
func (u *loginUsecase)Login(ctx context.Context,input *LoginInput)(*LoginOutput,error){
	output :=new(LoginOutput)
	data :=u.userService.LoginData(ctx,input.Email,input.Password)
	result:=u.userRepo.Login(ctx,data)
	output.Token = result
	return output,nil
}