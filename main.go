package main

import (
	"log"
	"vothings/configs"
	"vothings/internal/app/api"
	"vothings/internal/app/middleware"
	"vothings/internal/app/repository"
	"vothings/internal/app/service"
	"vothings/internal/pkg/driver"
	"vothings/internal/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := configs.LoadCfg()
	jwt := utils.NewJwtToken(cfg.IssuerName, cfg.JwtSignatureKey, cfg.JwtLifeTime)

	db, err := driver.GetConnectionPostgres(cfg)
	if err != nil {
		log.Fatalf("error when open connection postgres: %s", err)
	}

	candidateRepo := repository.NewCandidateRepository(db)
	voteRepo := repository.NewVoteRepository(db)
	userRepo := repository.NewRepository(db)

	service := service.NewUserService(userRepo, *jwt)

	middleware := middleware.NewAuthMiddleware(*jwt)

	app := fiber.New()
	api.NewAuth(app, service)
	api.NewCandidateApi(app, candidateRepo, *middleware)
	api.NewVoteAPi(app, voteRepo, *middleware)

	app.Listen(cfg.Server.Host + ":" + cfg.Server.Port)
}
