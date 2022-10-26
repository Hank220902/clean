package route

import (
	"clean/cmd/middle"


	"github.com/kataras/iris/v12"
	// "google.golang.org/grpc"
)

// var client pb.TodoListServiceClient

func Todo(app *iris.Application) {
	api := app.Party("/api")
	{
		api.Post("/todo",middle.Create)
		api.Get("/todo", middle.GetAll)
		// api.Post("/todolist", controllers.CreateToDoList)
		// api.Get("/manytodolist", controllers.GetAllToDoList)
		// api.Put("/todolist", controllers.UpdateToDoList)
		// api.Delete("/todolist", controllers.DeleteToDoList)

	}

}
