package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = []byte("3104bfa7b205eb55d6dc2c44f9185d44")

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateJWT(username string, password string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
