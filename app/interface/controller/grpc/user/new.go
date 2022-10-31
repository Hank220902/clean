package user

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/user/login"
	"clean/app/usecase/user/register"
)

type userServer struct {
	registerUsecase register.RegisterUsecase
	loginUsecase    login.LoginUsecase
}

func NewUserServer(registerUsecase register.RegisterUsecase,loginUsecase login.LoginUsecase) pb.UserServiceServer {
	return &userServer{
		registerUsecase: registerUsecase,
		loginUsecase: loginUsecase,
	}
}

