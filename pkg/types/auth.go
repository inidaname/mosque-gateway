package types

import "github.com/golang-jwt/jwt/v5"

type Claims = jwt.MapClaims
type Token = jwt.Token

type ApiKeyValidityCache struct {
	Valid bool
}

type Authenticator interface {
	GenerateToken(claims Claims) (string, error)
	ValidateToken(token string) (*Token, error)
}
