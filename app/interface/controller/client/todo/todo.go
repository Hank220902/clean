package client

import (
	"context"
	"fmt"
	"log"

	// "clean/app/interface/controller/grpc/todo"
	pb "clean/app/interface/controller/grpc/protobuf"

	"google.golang.org/grpc"
)



var todoClient pb.TodoServiceClient

func init() {
	connect, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	todoClient = pb.NewTodoServiceClient(connect)
}

func Create(ctx context.Context, in *CreateInput) int32 {
	result, err := todoClient.Create(ctx, &pb.CreateRequest{
		Matter:            in.Matter,
		EndTime:           in.EndTime,
		FinishedCondition: in.FinishedCondition,
		Email:             in.Email,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.GetResMessage())
	return result.GetResMessage()
	// return result.GetResMessage()
}

func GetAll(ctx context.Context, in *GetAllInput)*pb.GetResponse{
	result,err := todoClient.GetAll(ctx,&pb.GetRequest{
		Email: in.Email,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.GetGetResult())
	return result
}

func Delete(ctx context.Context, in *DeleteInput)*pb.DeleteResponse{
	result,err :=todoClient.Delete(ctx,&pb.DeleteRequest{
		Id: in.Id,
		Email: in.Email,
	})
	if err!= nil {
        panic(err)
    }
	return result
}

func Update(ctx context.Context, in *UpdateInput)*pb.UpdateResponse{
	result,err :=todoClient.Update(ctx,&pb.UpdateRequest{
		Id: in.Id,
		Email: in.Email,
		FinishedCondition: in.FinishedCondition,
		Note: in.Note,
	})
	if err!= nil {
        panic(err)
    }
	return result
}


