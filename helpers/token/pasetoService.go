package token

import (
	"fmt"
	"os"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

const MaxAge = time.Minute * 5

func getSecretKey() string {
	secret := os.Getenv("PASETO_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker() (Maker, error) {
	symmetricKey := getSecretKey()
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) GenerateToken(name string, password string) string {
	payload := paseto.JSONToken{
		Issuer:     "go-ifsc",
		Expiration: time.Now().Add(MaxAge),
		IssuedAt:   time.Now(),
	}
	payload.Set("user", name)
	payload.Set("password", password)

	t, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		panic(err)
	}
	return t
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) ValidateToken(tokenString string) (*paseto.JSONToken, error) {
	jsonToken := paseto.JSONToken{}

	err := maker.paseto.Decrypt(tokenString, maker.symmetricKey, &jsonToken, nil)
	if err != nil {
		return nil, err
	}
	err = jsonToken.Validate()
	if err != nil {
		return nil, err
	}
	return &jsonToken, nil
}
