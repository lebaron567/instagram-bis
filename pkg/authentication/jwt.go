package authentication

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

// Clé secrète pour signer le JWT
var jwtKey = []byte("c8f9d72e3b4a6d9e7f0b1c2a3e4f5g6h7i8j9k0l1m2n3o4p5q6r7s8t9u0v1w2x3")

func CreateToken(email string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Le token expire après 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
