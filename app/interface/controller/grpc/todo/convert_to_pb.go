package todo

import (
	
	pb "clean/app/interface/controller/grpc/protobuf"
	// "clean/app/domain/model/todo"
	"clean/app/usecase/todo/getall"
)

func convertGetTodosToPb(in []*getall.GetAllOutput) []*pb.GetResult {
	ans := make([]*pb.GetResult, len(in))
	for i, v := range in {
		ans[i] = convertGetTodoToPb(v)

	}
	return ans
}

func convertGetTodoToPb(in *getall.GetAllOutput) *pb.GetResult {
	if in == nil {
		return new(pb.GetResult)
	}
	return &pb.GetResult{
		Id:                in.ID,
		Matter:            in.Matter,
		EndTime:           in.EndTime,
		FinishedCondition: in.FinishedCondition,
		CreateTime:        in.CreateTime,
		Status:            in.Status,
		Note:              in.Note,
		Email:             in.Email,
		UpdateTime:        in.UpdateTime,
	}
}
