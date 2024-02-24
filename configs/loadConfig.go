package configs

import (
	"log"
	"os"
	"strconv"
	"time"
	"vothings/internal/app/domain"

	"github.com/joho/godotenv"
)

func LoadCfg() *domain.Configs {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error when load configs: %s", err)
	}

	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))

	if err != nil {
		log.Fatalf("error when load configs: %s", err)
	}

	return &domain.Configs{
		Server: domain.Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Postgres: domain.Postgres{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			UserName: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
		JwtToken: domain.JwtToken{
			JwtLifeTime:     time.Duration(tokenLifeTime) * time.Hour,
			JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
			IssuerName:      os.Getenv("TOKEN_ISSUE_NAME"),
		},
	}
}
