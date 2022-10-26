package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"context"
	"fmt"
	"net"

	ser "clean/app/domain/service/todo"
	repo "clean/app/infra/presistence/mongo/todo"
	"clean/app/usecase/todo/create"

	"google.golang.org/grpc"
)

const (
	success     int = 1
	fail        int = 2
	tokenError  int = 3
	emailExists int = 4
)

type todoService struct {
	pb.UnimplementedTodoServiceServer
}

func (s *server) Create(ctx context.Context, data *pb.CreateRequest) (*pb.CreateResponse, error) {
	fmt.Println("qqqqqasdadadasd")
	fmt.Println(data)
	fmt.Println("qqqqqasdad12313")
	fmt.Println(convertToCreateInput(data))
	fmt.Println(s.createUsecase)
	_, err := s.createUsecase.Create(ctx, convertToCreateInput(data))
	fmt.Println("aaaaaa")
	if err != nil {
		return new(pb.CreateResponse), err
	}

	return &pb.CreateResponse{ResMessage: int32(success)}, nil
}

func GrpcServer() {
	rpcs := grpc.NewServer()

	NewRepo := repo.NewRepository()
	NewService := ser.NewTodoService()
	NewCreateUsecase := create.NewCreateUsecase(NewRepo, NewService)
	pb.RegisterTodoServiceServer(rpcs, NewServer(NewCreateUsecase))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
