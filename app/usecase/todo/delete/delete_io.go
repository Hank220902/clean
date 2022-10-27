package delete


type DeleteInput struct {
	Id string `json:"id"`
	Email string `json:"email"`
}

type DeleteOutput struct {
	Message int
}
