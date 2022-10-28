package middle

import (
	"clean/app/infra/encryption/bcrypt"
	client "clean/app/interface/controller/client/user"
	"fmt"
	"clean/app/infra/token/jwt"

	"github.com/kataras/iris/v12"
)

func Register(ctx iris.Context) {
	var RegisterData client.RegisterInput
	if err := ctx.ReadJSON(&RegisterData); err != nil {
		panic(err.Error())
	}

	hashStr, err := bcrypt.HashAndSalt(RegisterData.Password)
	if err != nil {
		fmt.Println(err)
	}


	input := client.RegisterInput{
		Name:     RegisterData.Name,
		Email:    RegisterData.Email,
		Password: hashStr,
	}

	result := client.Register(requestContext(ctx), &input)

	ctx.JSON(result)
}

func Login(ctx iris.Context){
	var LoginData client.LoginInput
	if err := ctx.ReadJSON(&LoginData); err != nil {
		panic(err.Error())
	}
	input := client.LoginInput{
		Email: LoginData.Email,
		Password: LoginData.Password,
	}
	result := client.Login(requestContext(ctx), &input)
	ctx.JSON(result)
}

func Logout(ctx iris.Context){
	email := jwt.MyAuthenticatedHandler(ctx)
	result := jwt.DeleteToken(ctx,email)
	ctx.JSON(result)
}


