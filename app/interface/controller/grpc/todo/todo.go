package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"context"
	"net"

	service "clean/app/domain/service/todo"
	repo "clean/app/infra/presistence/mongo/todo"
	"clean/app/usecase/todo/create"
	"clean/app/usecase/todo/getall"

	"google.golang.org/grpc"
)

const (
	success     int = 1
	fail        int = 2
	tokenError  int = 3
	emailExists int = 4
)



func (s *server) Create(ctx context.Context, data *pb.CreateRequest) (*pb.CreateResponse, error) {

	_, err := s.createUsecase.Create(ctx, convertToCreateInput(data))

	if err != nil {
		return new(pb.CreateResponse), err
	}

	return &pb.CreateResponse{ResMessage: int32(success)}, nil
}

func (s *server) GetAll(ctx context.Context, data *pb.GetRequest) (*pb.GetResponse,error){
	result,err :=s.getAllUsecase.GetAll(ctx,convertToGetAllInput(data))
	if err!= nil {
        return new(pb.GetResponse),err
	}
	return &pb.GetResponse{GetResult: convertGetTodosToPb(result)},nil
}

func GrpcServer() {
	rpcs := grpc.NewServer()

	NewRepo := repo.NewRepository()
	NewService := service.NewTodoService()
	NewCreateUsecase := create.NewCreateUsecase(NewRepo, NewService)
	NewGetAllUsecase := getall.NewGetAllUsecase(NewRepo, NewService)
	pb.RegisterTodoServiceServer(rpcs, NewServer(NewCreateUsecase,NewGetAllUsecase))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
