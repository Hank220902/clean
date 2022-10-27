package service

import (
	model "clean/app/domain/model/todo"
	"context"
	"fmt"
	"time"
)

func (u *todoService) CreateData(ctx context.Context, Matter, EndTime, FinishedCondition, Email string) (*model.Create, error) {
	var CreateData = new(model.Create)
	CreateData.Matter = Matter
	CreateData.CreateTime = time.Now()
	emdTimeResult, err := time.Parse(time.RFC3339, EndTime)
	if err != nil {
		fmt.Println(err)
	}
	CreateData.EndTime = emdTimeResult
	CreateData.FinishedCondition = FinishedCondition
	CreateData.Email = Email
	CreateData.Status = ""
	return CreateData, nil
}

func(u *todoService) DeleteData(ctx context.Context, Email string,Id string)(*model.Delete){
	var delete = new(model.Delete)
    delete.Email = Email
    delete.ID = Id
    return delete
}

func (u *todoService)UpdateData(ctx context.Context, Email,Id,FinishedCondition,Note string)(*model.Update){
	var update = new(model.Update)
	update.Email=Email
	update.FinishedCondition= FinishedCondition
	update.ID= Id
	update.Note= Note
	return update
}
