package model

import "time"

type Create struct {
	Matter            string    `json:"matter" bson:"matter"`
	EndTime           time.Time `json:"endTime" bson:"endTime"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`
	Status            string    `json:"status" bson:"status"`
	Email             string    `json:"email" bson:"email"`
	CreateTime        time.Time `json:"createTime" bson:"createTime"`
}

type Get struct {
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

type GetFilter struct {
	Email             string    `json:"email" bson:"email"`
	Status            string    `json:"status" bson:"status"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`

}

type UpdateStatus struct {
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

type Delete struct {
	ID    string `json:"_id" bson:"_id"`
	Email string `json:"email" bson:"email"`
}

type Update struct {
	ID string `json:"_id" bson:"_id"`
	Email string `json:"email" bson:"email"`
	FinishedCondition string    `json:"finishedCondition" bson:"finishedCondition"`
	UpdateTime        time.Time `json:"updateTime" bson:"updateTime"`
	Note              string    `json:"note" bson:"note"`
	Status            string    `json:"status" bson:"status"`
}
