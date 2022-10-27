package middle

import (
	client "clean/app/interface/controller/client/todo"

	"context"

	"github.com/kataras/iris/v12"
)

func Create(ctx iris.Context) {

	var Todo client.CreateInput
	if err := ctx.ReadJSON(&Todo); err != nil {
		panic(err.Error())
	}

	input := client.CreateInput{
		Matter:            Todo.Matter,
		EndTime:           Todo.EndTime,
		FinishedCondition: Todo.FinishedCondition,
		Email:             Todo.Email,
		Status:            "",
	}

	result := client.Create(requestContext(ctx), &input)

	ctx.JSON(result)

}
func requestContext(ctx iris.Context) context.Context {
	return ctx.Request().Context()
}

func GetAll(ctx iris.Context) {

	paramsEmail := ctx.URLParam("email")
	input := client.GetAllInput{
		Email: paramsEmail,
	}

	result := client.GetAll(requestContext(ctx), &input)

	ctx.JSON(result)
}

func Delete(ctx iris.Context) {
	var Input client.DeleteInput
	if err := ctx.ReadJSON(&Input); err != nil {
		panic(err.Error())
	}
	paramsId := ctx.URLParam("id")
	data := client.DeleteInput{
		Id:    paramsId,
		Email: Input.Email,
	}

	result := client.Delete(requestContext(ctx), &data)
	ctx.JSON(result.ResMessage)
}

func Update(ctx iris.Context) {
	var Input client.UpdateInput
	if err := ctx.ReadJSON(&Input); err != nil {
		panic(err.Error())
	}
	data := client.UpdateInput{
		Id:                Input.Id,
		Email:             Input.Email,
		FinishedCondition: Input.FinishedCondition,
		Note:              Input.Note,
	}
	result := client.Update(requestContext(ctx), &data)
	ctx.JSON(result)
}
