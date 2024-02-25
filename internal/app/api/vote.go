package api

import (
	"vothings/internal/app/dto"
	"vothings/internal/app/middleware"
	"vothings/internal/app/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoteRepository interface {
	Insert(vote model.Votes) error
}

type votesApi struct {
	vote       VoteRepository
	middleware middleware.AuthMiddleware
}

func (v *votesApi) CreateVotes(ctx *fiber.Ctx) error {
	var votesData dto.VoteDtoRequest

	if err := ctx.BodyParser(&votesData); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "failed to read data"})
	}

	candidateId, err := uuid.Parse(votesData.CandidateId)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "failed to read data"})
	}

	userId, err := uuid.Parse(votesData.UserId)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "failed to read data"})
	}

	if err := v.vote.Insert(model.Votes{
		CandidateId: candidateId,
		UserID:      userId,
	}); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "vote created successfuly"})
}

func NewVoteAPi(App *fiber.App, vote VoteRepository, middleware middleware.AuthMiddleware) {
	h := votesApi{
		vote:       vote,
		middleware: middleware,
	}

	App.Post("vote/create", h.middleware.RequireToken("user"), h.CreateVotes)
}
