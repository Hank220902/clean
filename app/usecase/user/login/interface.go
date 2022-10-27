package login

import "context"

type LoginUsecase interface {
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
}
