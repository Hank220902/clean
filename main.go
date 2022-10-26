package main

import (
	mon "clean/app/infra/presistence/mongo"
	"clean/app/interface/controller/grpc/todo"
	// "clean/cmd"
	"clean/cmd/route"
	"fmt"

	"github.com/kataras/iris/v12"
)

func init() {
	go todo.GrpcServer()
	// go cmd.Server()
	mon.Connect()
}

func main() {

	app := iris.New()
	fmt.Println("hello")

	route.Todo(app)

	//cmd.Execute()

	app.Listen(":3000")

}
