package token

import "github.com/o1egl/paseto"

type Maker interface {
	GenerateToken(name string, password string) string
	ValidateToken(tokenString string) (*paseto.JSONToken, error)
}
