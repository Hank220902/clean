package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"clean/cmd"
	"clean/app/interface/controller/grpc/todo"
)

func init() {
	go todo.GrpcServer()
	


}

func main() {

	app := iris.New()
	fmt.Println("hello")

	route.Todo(app)


	app.Listen(":3000")

}
