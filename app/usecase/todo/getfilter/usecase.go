package getfilter

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

type getFilterUsecase struct {
	todoService service.TodoService
	todoRepo    repository.TodoRepository
}

func NewGetFilterUsecase(todoRepo repository.TodoRepository, todoService service.TodoService) GetFilterUsecase {
	return &getFilterUsecase{
		todoRepo:    todoRepo,
		todoService: todoService,
	}
}

func (u *getFilterUsecase) GetFilter(ctx context.Context, input *GetFilterInput) ([]*GetFilterOutput, error) {
	// output := new(*[]GetAllOutput)
	data:= u.todoService.GetfilterData(ctx,input.Email, input.Status, input.FinishedCondition)

	list, err := u.todoRepo.GetFilter(ctx, data)
	if err != nil {
		return nil, err
	}

	output := convertGetFilterListToOutput(list)

	return output, nil
}

