package token

import "github.com/golang-jwt/jwt"

type Maker interface {
	GenerateToken(name string, password string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}
