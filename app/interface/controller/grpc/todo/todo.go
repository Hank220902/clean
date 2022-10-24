package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"context"
)

type todoService struct {
	pb.UnimplementedTodoServiceServer
}

func (ts *todoService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error){
	
}
