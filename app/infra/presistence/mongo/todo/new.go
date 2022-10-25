package todo

import (
	repository "clean/app/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// repo -
type repo struct {
	db *mongo.Client
}

// NewRepository -
func NewRepository(m *mongo.Client) repository.TodoRepository {
	return &repo{
		db: m,
	}
}
