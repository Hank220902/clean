package todo

import (
	repository "clean/app/domain/repository"
)

// repo -
type repo struct {
}

// NewRepository -
func NewRepository() repository.TodoRepository {
	return &repo{}
}
