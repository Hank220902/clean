package getall

import (

	repository "clean/app/domain/repository"
	service "clean/app/domain/service/todo"

	

	"context"
)

const (
	success     int = 1
	fail        int = 2
	tokenError  int = 3
	emailExists int = 4
)

type getAllUsecase struct {
	todoService service.TodoService
	todoRepo    repository.TodoRepository
}

func NewGetAllUsecase(todoRepo repository.TodoRepository, todoService service.TodoService) GetAllUsecase {
	return &getAllUsecase{
		todoRepo:    todoRepo,
		todoService: todoService,
	}
}

func (u *getAllUsecase) GetAll(ctx context.Context, input *GetAllInput) ([]*GetAllOutput, error) {
	// output := new(*[]GetAllOutput)
	list, err := u.todoRepo.GetAll(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	output := convertGetAllListToOutput(list)

	return output, nil
}

