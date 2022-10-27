package middle

import (
	"clean/app/infra/encryption/bcrypt"
	client "clean/app/interface/controller/client/user"
	"fmt"

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
