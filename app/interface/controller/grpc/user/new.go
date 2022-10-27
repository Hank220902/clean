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

// import (
// 	pb "clean/app/interface/controller/grpc/protobuf"
// 	"clean/app/usecase/todo/create"
// 	"clean/app/usecase/todo/getall"
// 	"clean/app/usecase/todo/delete"
// 	"clean/app/usecase/todo/update"
// )

// type server struct {
// 	createUsecase create.CreateUsecase
// 	getAllUsecase getall.GetAllUsecase
// 	deleteUsecase delete.DeleteUsecase
// 	updateUsecase update.UpdateUsecase
// }

// func NewServer(createUsecase create.CreateUsecase,getAllUsecase getall.GetAllUsecase,deleteUsecase delete.DeleteUsecase,updateUsecase update.UpdateUsecase) pb.TodoServiceServer {
// 	return &server{
// 		createUsecase: createUsecase,
// 		getAllUsecase: getAllUsecase,
// 		deleteUsecase: deleteUsecase,
// 		updateUsecase: updateUsecase,
// 	}
// }
