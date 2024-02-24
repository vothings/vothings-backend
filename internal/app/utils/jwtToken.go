package utils

import (
	"fmt"
	"vothings/internal/app/domain"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	cfg *domain.Configs
}

func (j *Jwt) GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      j.cfg.JwtLifeTime,
	})
	tokenString, err := token.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *Jwt) VerivyToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
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

func NewJwtToken(cfg *domain.Configs) *Jwt {
	return &Jwt{
		cfg: cfg,
	}
}
