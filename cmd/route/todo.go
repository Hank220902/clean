package route

import (
	"clean/cmd/middle"

	"clean/app/infra/token/jwt"

	"github.com/kataras/iris/v12"
	// "google.golang.org/grpc"
)

// var client pb.TodoListServiceClient

func Todo(app *iris.Application) {
	api := app.Party("/api", jwt.J.Serve)
	{
		api.Post("/todo", middle.Create)
		api.Get("/todo", middle.GetAll)
		api.Delete("/todo", middle.Delete)
		api.Put("/todo", middle.Update)


	}

}
