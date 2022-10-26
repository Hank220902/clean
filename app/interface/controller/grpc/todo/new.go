package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
	"clean/app/usecase/todo/getall"
	"clean/app/usecase/todo/delete"
)

type server struct {
	createUsecase create.CreateUsecase
	getAllUsecase getall.GetAllUsecase
	deleteUsecase delete.DeleteUsecase
}

func NewServer(createUsecase create.CreateUsecase,getAllUsecase getall.GetAllUsecase,deleteUsecase delete.DeleteUsecase) pb.TodoServiceServer {
	return &server{
		createUsecase: createUsecase,
		getAllUsecase: getAllUsecase,
		deleteUsecase: deleteUsecase,
	}
}
