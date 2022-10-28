package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
	"clean/app/usecase/todo/getall"
	"clean/app/usecase/todo/delete"
	"clean/app/usecase/todo/update"
	"clean/app/usecase/todo/getfilter"
)

type todoServer struct {
	createUsecase create.CreateUsecase
	getAllUsecase getall.GetAllUsecase
	getFilterUsecase getfilter.GetFilterUsecase
	deleteUsecase delete.DeleteUsecase
	updateUsecase update.UpdateUsecase

}

func NewTodoServer(createUsecase create.CreateUsecase,getAllUsecase getall.GetAllUsecase,deleteUsecase delete.DeleteUsecase,updateUsecase update.UpdateUsecase,getFilterUsecase getfilter.GetFilterUsecase) pb.TodoServiceServer {
	return &todoServer{
		createUsecase: createUsecase,
		getAllUsecase: getAllUsecase,
		deleteUsecase: deleteUsecase,
		updateUsecase: updateUsecase,
		getFilterUsecase: getFilterUsecase,
	}
}
