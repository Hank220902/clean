package update

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

type updateUsecase struct {
	todoService service.TodoService
	todoRepo    repository.TodoRepository
}

func NewUpdateUsecase(todoRepo repository.TodoRepository, todoService service.TodoService) UpdateUsecase {
	return &updateUsecase{
		todoRepo:    todoRepo,
		todoService: todoService,
	}
}

func (u *updateUsecase) Update(ctx context.Context, Input *UpdateInput) (*UpdateOutput, error) {
	output := new(UpdateOutput)

	data := u.todoService.UpdateData(ctx, Input.Email, Input.Id, Input.FinishedCondition, Input.Note)

	result := u.todoRepo.Update(ctx, data)
	output.Message = result
	return output, nil

}
