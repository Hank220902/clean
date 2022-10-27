package grpc

import (
	todoService "clean/app/domain/service/todo"
	userService "clean/app/domain/service/user"
	todoRepo "clean/app/infra/presistence/mongo/todo"
	userRepo "clean/app/infra/presistence/mongo/user"
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
	"clean/app/usecase/todo/delete"
	"clean/app/usecase/todo/getall"
	"clean/app/usecase/todo/update"
	"clean/app/usecase/user/register"
	"net"
	"clean/app/interface/controller/grpc/todo"
	"clean/app/interface/controller/grpc/user"

	"google.golang.org/grpc"
)

func GrpcServer() {
	rpcs := grpc.NewServer()

	NewRepo := todoRepo.NewRepository()
	NewService := todoService.NewTodoService()
	NewUserRepo := userRepo.NewRepository()
	NewUserService := userService.NewUserService()
	NewRegisterUsecase := register.NewRegisterUsecase(NewUserRepo, NewUserService)
	NewCreateUsecase := create.NewCreateUsecase(NewRepo, NewService)
	NewGetAllUsecase := getall.NewGetAllUsecase(NewRepo, NewService)
	NewDeleteUsecase := delete.NewDeleteUsecase(NewRepo, NewService)
	NewUpdaeteUsecase := update.NewUpdateUsecase(NewRepo, NewService)
	pb.RegisterTodoServiceServer(rpcs, todo.NewTodoServer(NewCreateUsecase, NewGetAllUsecase, NewDeleteUsecase, NewUpdaeteUsecase))
	pb.RegisterUserServiceServer(rpcs, user.NewUserServer(NewRegisterUsecase))
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
