package api

import (
	"vothings/internal/app/dto"
	"vothings/internal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type VotesService interface {
	CreateVote(voteData dto.VoteDtoRequest) error
}

type votesApi struct {
	service    VotesService
	middleware middleware.AuthMiddleware
}

func (v *votesApi) CreateVotes(ctx *fiber.Ctx) error {
	var votesData dto.VoteDtoRequest

	if err := ctx.BodyParser(&votesData); err != nil {
		return ctx.SendStatus(400)
	}

	if err := v.service.CreateVote(votesData); err != nil {
		return ctx.SendStatus(400)
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "vote created successfuly"})
}

func NewVoteAPi(App *fiber.App, service VotesService, middleware middleware.AuthMiddleware) {
	h := votesApi{
		service:    service,
		middleware: middleware,
	}

	App.Post("vote/create", h.middleware.RequireToken("user"), h.CreateVotes)
}
