package register

import "context"

type RegisterUsecase interface {
	Register(ctx context.Context, input *RegisterInput) (*RegisterOutput, error)
}
