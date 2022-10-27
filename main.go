package main

import (
	"clean/app/infra/presistence/mongo"
	"clean/app/infra/memory/redis"
	"clean/app/interface/controller/grpc"
	// "clean/cmd"
	"clean/cmd/route"
	"fmt"

	"github.com/kataras/iris/v12"
)

func init() {
	go grpc.GrpcServer()
	// go cmd.Server()
	mongo.Connect()
	redis.Redis()
	
}

func main() {

	app := iris.New()
	fmt.Println("hello")

	route.Todo(app)
	route.User(app)
	
	//cmd.Execute()

	app.Listen(":3000")

}
