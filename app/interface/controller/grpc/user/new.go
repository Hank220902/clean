package user

import ("clean/app/usecase/user/register"
pb "clean/app/interface/controller/grpc/protobuf"
)

type userServer struct {
	registerUsecase register.RegisterUsecase
}

func NewUserServer(registerUsecase register.RegisterUsecase) pb.UserServiceServer {
	return &userServer{
		registerUsecase: registerUsecase,
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
