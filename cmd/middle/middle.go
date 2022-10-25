package middle

import (
	"clean/app/interface/controller/client"

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

	result := client.Create(RequestContext(ctx), &input)

	ctx.JSON(result)

}
func RequestContext(ctx iris.Context) context.Context {
	return ctx.Request().Context()
}
