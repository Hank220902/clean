package delete


type DeleteInput struct {
	ID string `json:"id"`
	Email string `json:"email"`
}

type DeleteOutput struct {
	Message int
}
