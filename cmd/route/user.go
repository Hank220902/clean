package route

import (
	"clean/app/infra/token/jwt"
	"clean/cmd/middle"

	"github.com/kataras/iris/v12"
)

func User(app *iris.Application) {
	user := app.Party("/user")
	{
		user.Post("/register", middle.Register)
		user.Post("/login", middle.Login)
		// api.Get("/todo", middle.GetAll)
		// api.Delete("/todo", middle.Delete)
		// api.Put("/todo", middle.Update)
		// api.Post("/todolist", controllers.CreateToDoList)
		// api.Get("/manytodolist", controllers.GetAllToDoList)
		// api.Put("/todolist", controllers.UpdateToDoList)
		// api.Delete("/todolist", controllers.DeleteToDoList)
		user.Get("/logout",jwt.J.Serve,middle.Logout)

	}
	// user = app.Party("/logout", jwt.J.Serve)
	// {
	// 	user.Get("", middle.Logout)

	// }

}
