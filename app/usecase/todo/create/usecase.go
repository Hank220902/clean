package create

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

type createUsecase struct {
	todoService service.TodoService
	todoRepo    repository.TodoRepository
}

func NewCreateUsecase(todoRepo repository.TodoRepository, todoService service.TodoService) CreateUsecase {
	return &createUsecase{
		todoRepo:    todoRepo,
		todoService: todoService,
	}
}

func (u *createUsecase) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {

	output := new(CreateOutput)

	todo, err := u.todoService.NewTodo(ctx, input.Matter, input.EndTime, input.FinishedCondition, input.Email)

	if err != nil {
		return nil, err
	}

	if err := u.todoRepo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return output, nil
}
