package getfilter

type GetFilterInput struct {
	Email string `json:"email" bson:"email"`
	Status string `json:"status" bson:"status"`
	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`
}

type GetFilterOutput struct {
	ID                string `json:"_id" bson:"_id"`
	Matter            string `json:"matter" bson:"matter"`
	EndTime           string `json:"endTime" bson:"endTime"`
	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`
	Status            string `json:"status" bson:"status"`
	Email             string `json:"email" bson:"email"`
	Note              string `json:"note" bson:"note"`
	CreateTime        string `json:"createTime" bson:"createTime"`
	UpdateTime        string `json:"updateTime" bson:"updateTime"`
}

// type GetAllOutput struct {
// 	List []*GetAllList
// }
