package user

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"shopping-service/driver"
	"shopping-service/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"shopping-service/constants"
)

type User interface {
	Login(string, string) (string, error)
}

type users struct {
	secret   string
	collection string
	executor driver.MongoConnector
}

func NewUser(coll, secret string, e driver.MongoConnector) User {
	return &users{collection: coll, secret: secret, executor: e}
}

func (u *users) Login(username, password string) (string, error) {
	result, err := u.executor.Reader(bson.D{{"username", username}, {"password", password}}, u.collection)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve user credentials: %s", err)
	}

	if result == nil {
		return "", fmt.Errorf("invalid credentials")
	}

	secret := []byte(u.secret)
	claims := jwt.MapClaims{}
	claims["role"] = result.(models.Result)["role"]
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("unable to sign token: %s", err)
	}

	return fmt.Sprintf("%s%s", constants.Bearer, signedToken), nil
}