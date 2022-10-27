package user

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/user/register"
)

func convertToRegisterInput(in *pb.RegisterRequest) *register.RegisterInput {
	return &register.RegisterInput{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}
