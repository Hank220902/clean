package todo

import (
	pb "clean/app/interface/controller/grpc/protobuf"
	"context"
)

const (
	success     int = 1
	fail        int = 2
	tokenError  int = 3
	emailExists int = 4
)

func (s *todoServer) Create(ctx context.Context, data *pb.CreateRequest) (*pb.CreateResponse, error) {

	_, err := s.createUsecase.Create(ctx, convertToCreateInput(data))

	if err != nil {
		return new(pb.CreateResponse), err
	}

	return &pb.CreateResponse{ResMessage: int32(success)}, nil
}

func (s *todoServer) GetAll(ctx context.Context, data *pb.GetRequest) (*pb.GetResponse, error) {
	result, err := s.getAllUsecase.GetAll(ctx, convertToGetAllInput(data))
	if err != nil {
		return new(pb.GetResponse), err
	}
	return &pb.GetResponse{GetResult: convertGetTodosToPb(result)}, nil
}

func (s *todoServer) GetFilter(ctx context.Context, data *pb.GetFilterRequest)(*pb.GetFilterResponse, error){
	result, err := s.getFilterUsecase.GetFilter(ctx,convertToGetFilterInput(data))
	if err != nil {
		return new(pb.GetFilterResponse), err
	}
	return &pb.GetFilterResponse{GetResult: convertGetFilterTodosToPb(result)},nil
}

func (s *todoServer) Delete(ctx context.Context, data *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result, err := s.deleteUsecase.Delete(ctx, convertToDeleteInput(data))
	if err != nil {
		return new(pb.DeleteResponse), err
	}
	return &pb.DeleteResponse{ResMessage: int32(result.Message)}, nil
}

func (s *todoServer) Update(ctx context.Context, data *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	result, err := s.updateUsecase.Update(ctx, convertToUpdateInput(data))

	if err != nil {
		return new(pb.UpdateResponse), err
	}
	return &pb.UpdateResponse{ResMessage: int32(result.Message)}, err
}
