package getfilter

import "context"

type GetFilterUsecase interface {
	GetFilter(ctx context.Context, input *GetFilterInput)([]*GetFilterOutput,error)
}
