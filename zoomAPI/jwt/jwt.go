package jwt_create

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT() (string, error) {

	mySigningKey := []byte("F65tiTOVh1HWT38Tc8dZZunvKa2PGsK2")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["iss"] = "RagRtzP8TLyIbjzYVGIsFQ"
	claims["exp"] = time.Now().Add(15 * time.Minute)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
