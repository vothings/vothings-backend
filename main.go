package main

import (
	"vothings/configs"
	"vothings/internal/app/api"
	"vothings/internal/app/driver"
	"vothings/internal/app/middleware"
	"vothings/internal/app/repository"
	"vothings/internal/app/service"
	"vothings/internal/app/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := configs.LoadCfg()
	jwt := utils.NewJwtToken(cfg)
	db := driver.GetConnectionPostgres(cfg)

	candidateRepo := repository.NewCandidateRepository(db)
	voteRepo := repository.NewVoteRepository(db)
	repo := repository.NewRepository(db)

	serviceCandidate := service.NewCandidateService(candidateRepo)
	serviceVote := service.NewVoteService(voteRepo)
	service := service.NewUserService(repo, *jwt)

	middleware := middleware.NewAuthMiddleware(*jwt)

	app := fiber.New()
	api.NewAuth(app, service)
	api.NewCandidateApi(app, &serviceCandidate, *middleware)
	api.NewVoteAPi(app, serviceVote, *middleware)

	app.Listen(cfg.Server.Host + ":" + cfg.Server.Port)
}
