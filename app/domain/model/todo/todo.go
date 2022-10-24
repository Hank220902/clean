package model

import "time"

type Todo struct {
	Matter            string    `json:"matter" bson:"matter"`
	EndTime           time.Time `json:"endTime" bson:"endTime"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`
	Status            string    `json:"status" bson:"status"`
	Email             string    `json:"email" bson:"email"`
	CreateTime        time.Time `json:"createTime" bson:"createTime"`
}

type HaveIDTodo struct {
	ID                string    `json:"_id" bson:"_id"`
	Matter            string    `json:"matter" bson:"matter"`
	EndTime           time.Time `json:"endTime" bson:"endTime"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`
	Status            string    `json:"status" bson:"status"`
	Email             string    `json:"email" bson:"email"`
	Note              string    `json:"note" bson:"note"`
	CreateTime        time.Time `json:"createTime" bson:"createTime"`
	UpdateTime        time.Time `json:"updateTime" bson:"updateTime"`
}
