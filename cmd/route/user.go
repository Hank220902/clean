package route

import (
	"clean/cmd/middle"

	"github.com/kataras/iris/v12"
)

func User(app *iris.Application) {
	api := app.Party("/user")
	{
		api.Post("/register", middle.Register)
		// api.Get("/todo", middle.GetAll)
		// api.Delete("/todo", middle.Delete)
		// api.Put("/todo", middle.Update)
		// api.Post("/todolist", controllers.CreateToDoList)
		// api.Get("/manytodolist", controllers.GetAllToDoList)
		// api.Put("/todolist", controllers.UpdateToDoList)
		// api.Delete("/todolist", controllers.DeleteToDoList)

	}

}
