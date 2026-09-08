package jwthelp

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetSecret(t *jwt.Token) (any, error) {
	return []byte(os.Getenv("JWT_SECRET")),nil
}