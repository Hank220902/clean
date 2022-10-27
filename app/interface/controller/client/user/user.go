package client

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"context"
	"log"

	"google.golang.org/grpc"
)

var userClient pb.UserServiceClient

func init() {
	connect, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	userClient = pb.NewUserServiceClient(connect)
}

func Register(ctx context.Context, in *RegisterInput) int32 {

	result, err := userClient.Register(ctx, &pb.RegisterRequest{
		Name:     in.Name,
		Password: in.Password,
		Email:    in.Email,
	})
	if err != nil {
		panic(err)
	}

	return result.GetResMessage()
}


