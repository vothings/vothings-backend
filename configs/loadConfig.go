package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Configs struct {
	Server
	Postgres
	JwtToken
}

type Server struct {
	Host string
	Port string
}

type Postgres struct {
	Port     string
	Host     string
	UserName string
	Password string
	DbName   string
}

type JwtToken struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
}

func LoadCfg() *Configs {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error when load configs: %s", err)
	}

	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))

	if err != nil {
		log.Fatalf("error when load configs: %s", err)
	}

	return &Configs{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Postgres: Postgres{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			UserName: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
		JwtToken: JwtToken{
			JwtLifeTime:     time.Duration(tokenLifeTime) * time.Hour,
			JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
			IssuerName:      os.Getenv("TOKEN_ISSUE_NAME"),
		},
	}
}
