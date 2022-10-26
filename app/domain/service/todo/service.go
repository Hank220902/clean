package service

import (
	model "clean/app/domain/model/todo"
	"context"
	"fmt"
	"time"
)

func (u *todoService) NewTodo(ctx context.Context, Matter, EndTime, FinishedCondition, Email string) (*model.Todo, error) {
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
	return newTodo, nil
}

// func (u *todoService) GetAll(ctx context.Context, Email string)(string){
// 	email := Email
// 	return email
// }
