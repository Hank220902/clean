package user

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"context"
)
const (
	success     int = 1
	fail        int = 2

)

func (s *userServer) Register(ctx context.Context, data *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	result, err := s.registerUsecase.Register(ctx,convertToRegisterInput(data))
	if err != nil {
		return new(pb.RegisterResponse),err
	}
	return &pb.RegisterResponse{ResMessage: int32(result.Message)}, nil
}

func(s *userServer) Login(ctx context.Context,data *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := s.loginUsecase.Login(ctx,convertToLoginInput(data))
	if err != nil {
		return new(pb.LoginResponse),err
	}
	return &pb.LoginResponse{ResMessage: result.Token},nil
}
