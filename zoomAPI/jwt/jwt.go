package jwt_create

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// func CreateJWT() (string, error) {

// 	mySigningKey := []byte("F65tiTOVh1HWT38Tc8dZZunvKa2PGsK2")

// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["iss"] = "RagRtzP8TLyIbjzYVGIsFQ"
// 	claims["exp"] = time.Now().Add(15 * time.Minute)

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func CreateJWT() (string, error) {
	var err error
	// here, we have kept it as 15 minute
	tokenExpirationTime := time.Now().Add(15 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	tokenClaims := jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: tokenExpirationTime.Unix(),
		Issuer:    "r8rWZb3hTEiMnh87x3Jz5A",
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, err := token.SignedString([]byte("V1liUm2rb2AfeQoVwLV0LLuxdL1Ch0eldD07"))

	return tokenString, err
}
