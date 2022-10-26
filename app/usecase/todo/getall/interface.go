package getall

import "context"

type GetAllUsecase interface {
	GetAll(ctx context.Context, input *GetAllInput)([]*GetAllOutput,error)
}
