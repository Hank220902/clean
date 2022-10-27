package client

type RegisterInput struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}