package user

import (
	repository "clean/app/domain/repository"
)

// repo -
type repo struct {
}

// NewRepository -
func NewRepository() repository.UserRepository {
	return &repo{}
}
