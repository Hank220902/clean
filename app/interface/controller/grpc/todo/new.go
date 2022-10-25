package todo

import(
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
)

type server struct {
	createUsecase create.CreateUsecase
}

func NewServer(createUsecase create.CreateUsecase) pb.TodoServiceServer {
	return &server{
		createUsecase:  createUsecase,

	}
}