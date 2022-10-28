package client

type CreateInput struct {
	Matter            string `json:"matter" bson:"matter"`
	EndTime           string `json:"endTime" bson:"endTime"`
	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`
	Status            string `json:"status" bson:"status"`
	Email             string `json:"email" bson:"email"`
}

type GetAllInput struct {
	Email string `json:"email" bson:"email"`
}

type GetFilterInput struct {
	Email string `json:"email" bson:"email"`
	Status string `json:"status" bson:"status"`
	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`
}

type DeleteInput struct {
	Id    string `json:"id" bson:"id"`
	Email string `json:"email" bson:"email"`
}

type UpdateInput struct {
	Id                string `json:"_id" bson:"_id"`
	Email             string `json:"email" bson:"email"`
	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`
	Note              string `json:"note" bson:"note"`
}
