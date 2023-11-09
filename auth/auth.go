package auth

import (
	"errors"
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
func ValidateJWT(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserName, claims.Password, nil
	}
	return "", "", errors.New("invalid token")
}
