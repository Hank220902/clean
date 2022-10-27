package update

type UpdateInput struct {
	Id string `json:"id"`

	FinishedCondition string `json:"finishedCondition" bson:"finishedCondition"`

	Email string `json:"email" bson:"email"`
	Note  string `json:"note" bson:"note"`
}

type UpdateOutput struct {
	Message int
}
