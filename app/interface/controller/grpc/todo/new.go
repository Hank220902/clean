package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
	"clean/app/usecase/todo/getall"
)

type server struct {
	createUsecase create.CreateUsecase
	getAllUsecase getall.GetAllUsecase
}

func NewServer(createUsecase create.CreateUsecase,getAllUsecase getall.GetAllUsecase) pb.TodoServiceServer {
	return &server{
		createUsecase: createUsecase,
		getAllUsecase: getAllUsecase,
	}
}
