package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"clean/app/usecase/todo/create"
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
