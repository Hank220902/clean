package delete

import (
	repository "clean/app/domain/repository"
	service "clean/app/domain/service/todo"
	"context"
)

type deleteUsecase struct {
	todoService service.TodoService
	todoRepo    repository.TodoRepository
}

func NewDeleteUsecase(todoRepo repository.TodoRepository, todoService service.TodoService) DeleteUsecase {
	return &deleteUsecase{
		todoRepo:    todoRepo,
		todoService: todoService,
	}
}

func (u *deleteUsecase) Delete(ctx context.Context, input *DeleteInput) (*DeleteOutput, error) {

	output := new(DeleteOutput)

	data := u.todoService.DeleteData(ctx, input.Email, input.Id)

	result := u.todoRepo.Delete(ctx, data)
	output.Message = result

	return output, nil
}
