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
		user.Get("/logout",jwt.J.Serve,middle.Logout)

	}


}
