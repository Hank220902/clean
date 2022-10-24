package service

import (
	model "clean/app/domain/model/todo"
	"context"
	"fmt"
	"time"
)

func NewTodo(ctx context.Context, Matter, EndTime, FinishedCondition,  Email string) (*model.Todo,error) {
	var newTodo = new(model.Todo)
    newTodo.Matter = Matter
	newTodo.CreateTime = time.Now()
	emdTimeResult, err := time.Parse(time.RFC3339, EndTime)
	if err != nil {
		fmt.Println(err)
	}
	newTodo.EndTime = emdTimeResult
	newTodo.FinishedCondition = FinishedCondition
	newTodo.Email = Email
    newTodo.Status = ""
	return newTodo,nil
}
