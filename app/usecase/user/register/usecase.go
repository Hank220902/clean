package register

import (
	repository "clean/app/domain/repository"
	service "clean/app/domain/service/user"
	"context"
)

const (
	success int = 1
	fail    int = 2
)

type registerUsecase struct {
	userService service.UserService
	userRepo    repository.UserRepository
}

func NewRegisterUsecase(userRepo repository.UserRepository, userService service.UserService) RegisterUsecase {
	return &registerUsecase{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (u *registerUsecase) Register(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
	output := new(RegisterOutput)
	data := u.userService.RegisterData(ctx, input.Name, input.Password, input.Email)
	result := u.userRepo.Register(ctx, data)

	output.Message = result

	return output, nil
}
