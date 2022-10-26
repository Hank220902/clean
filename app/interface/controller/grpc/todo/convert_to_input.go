package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
	"clean/app/usecase/todo/getall"
	"clean/app/usecase/todo/delete"
	"time"
)

func convertToCreateInput(in *pb.CreateRequest) *create.CreateInput {
	return &create.CreateInput{
		Matter:            in.Matter,
		FinishedCondition: in.FinishedCondition,
		EndTime:           in.EndTime,
		Email:             in.Email,
		CreateTime:        time.Now().Format(time.RFC3339),
	}
}

func convertToGetAllInput(in *pb.GetRequest) *getall.GetAllInput {
	return &getall.GetAllInput{
		Email: in.Email,
	}
}

func convertToDeleteInput(in *pb.DeleteRequest) *delete.DeleteInput {
	return &delete.DeleteInput{
        ID: in.Id,
		Email:in.Email,
    }
}
