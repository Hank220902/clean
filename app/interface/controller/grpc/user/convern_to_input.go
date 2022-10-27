package user

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/user/register"
	"clean/app/usecase/user/login"
)

func convertToRegisterInput(in *pb.RegisterRequest) *register.RegisterInput {
	return &register.RegisterInput{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func convertToLoginInput(in *pb.LoginRequest) *login.LoginInput {
	return &login.LoginInput{
		Email: in.Email,
		Password: in.Password,
	}

}
