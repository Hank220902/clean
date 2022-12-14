package create

type CreateInput struct {
	Matter            string `json:"matter" bson:"matter"`
	EndTime           string `json:"endTime" bson:"endTime"`
	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`
	Status            string `json:"status" bson:"status"`
	Email             string `json:"email" bson:"email"`
	CreateTime        string `json:"createTime" bson:"createTime"`
}

type CreateOutput struct {
	Message int
}
