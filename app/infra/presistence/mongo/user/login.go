package user

import (
	model "clean/app/domain/model/user"
	"clean/app/infra/encryption/bcrypt"
	"clean/app/infra/presistence/mongo"
	"clean/app/infra/token/jwt"


	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Login(ctx context.Context, data *model.Login) string {
	var LoginData model.Login

	filter := bson.D{{Key: "email", Value: data.Email}}
	err := mongo.UserCollection.FindOne(ctx, filter).Decode(&LoginData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Email)

	// token := jwt.GetTokenHandler(data.Email)
	check := bcrypt.ComparePasswords(LoginData.Password, data.Password)
	if !check {
		return "fail"
	} else {
		return jwt.GetTokenHandler(data.Email)
	}
}
