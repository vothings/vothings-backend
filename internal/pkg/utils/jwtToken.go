package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
}

func (j *Jwt) GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      j.JwtLifeTime,
	})
	tokenString, err := token.SignedString(j.JwtSignatureKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *Jwt) VerivyToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// fmt.Println(claims["role"], "KEKGINI")
	return claims, nil
}

func NewJwtToken(IssuerName string, JwtSignatureKey []byte, JwtLifeTime time.Duration) *Jwt {
	return &Jwt{
		IssuerName:      IssuerName,
		JwtSignatureKey: JwtSignatureKey,
		JwtLifeTime:     JwtLifeTime,
	}
}
