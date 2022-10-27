package middle

import (
	client "clean/app/interface/controller/client/todo"
	"clean/app/infra/token/jwt"

	"context"

	"github.com/kataras/iris/v12"
)

const(
	success int = 1
	fail int =2
	tokenError int = 3
	emailExists int = 4
)

func Create(ctx iris.Context) {
	email := jwt.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		ctx.JSON(tokenError)
		return
	}

	var Todo client.CreateInput
	if err := ctx.ReadJSON(&Todo); err != nil {
		panic(err.Error())
	}

	input := client.CreateInput{
		Matter:            Todo.Matter,
		EndTime:           Todo.EndTime,
		FinishedCondition: Todo.FinishedCondition,
		Email:             email,
		Status:            "",
	}

	result := client.Create(requestContext(ctx), &input)

	ctx.JSON(result)

}
func requestContext(ctx iris.Context) context.Context {
	return ctx.Request().Context()
}

func GetAll(ctx iris.Context) {
	email := jwt.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		ctx.JSON(tokenError)
		return
	}

	input := client.GetAllInput{
		Email: email,
	}

	result := client.GetAll(requestContext(ctx), &input)

	ctx.JSON(result)
}

func Delete(ctx iris.Context) {
	email := jwt.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		ctx.JSON(tokenError)
		return
	}
	var Input client.DeleteInput
	if err := ctx.ReadJSON(&Input); err != nil {
		panic(err.Error())
	}
	paramsId := ctx.URLParam("id")
	data := client.DeleteInput{
		Id:    paramsId,
		Email: email,
	}

	result := client.Delete(requestContext(ctx), &data)
	ctx.JSON(result.ResMessage)
}

func Update(ctx iris.Context) {
	email := jwt.MyAuthenticatedHandler(ctx)
	if email == "token not found" {
		ctx.JSON(tokenError)
		return
	}
	var Input client.UpdateInput
	if err := ctx.ReadJSON(&Input); err != nil {
		panic(err.Error())
	}
	data := client.UpdateInput{
		Id:                Input.Id,
		Email:             email,
		FinishedCondition: Input.FinishedCondition,
		Note:              Input.Note,
	}
	result := client.Update(requestContext(ctx), &data)
	ctx.JSON(result)
}
