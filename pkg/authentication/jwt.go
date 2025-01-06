package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("c8f9d72e3b4a6d9e7f0b1c2a3e4f5g6h7i8j9k0l1m2n3o4p5q6r7s8t9u0v1w2x3")

type Claims struct {
	UserID int    `json:"user_id"`
	Pseudo string `json:"pseudo"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, pseudo string, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Pseudo: pseudo,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
