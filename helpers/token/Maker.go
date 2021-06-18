package token

import "github.com/dgrijalva/jwt-go"

type Maker interface {
	GenerateToken(name string, password string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}
