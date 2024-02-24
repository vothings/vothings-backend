package domain

import "time"

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
