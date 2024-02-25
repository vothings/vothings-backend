package api

import (
	"vothings/internal/app/dto"
	"vothings/internal/app/middleware"
	"vothings/internal/app/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CandidateService interface {
	CreateCandidate(user dto.CandidateDtoRequest) error
	GetAllCandidate() ([]model.Candidates, error)
	UpdateCandidate(id uuid.UUID, candidate dto.CandidateDtoRequest) error
	DeleteCandidate(id uuid.UUID) error
}

type candidateApi struct {
	service    CandidateService
	middleware middleware.AuthMiddleware
}

func (c *candidateApi) Createcandidate(ctx *fiber.Ctx) error {
	var candidate dto.CandidateDtoRequest

	if err := ctx.BodyParser(&candidate); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "error when parsing data"})
	}

	if err := c.service.CreateCandidate(candidate); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "candidate created successfuly"})
}

func (c *candidateApi) GetAllcandidate(ctx *fiber.Ctx) error {
	candidates, err := c.service.GetAllCandidate()

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "error when parsing data"})
	}

	return ctx.Status(200).JSON(candidates)
}

func (c *candidateApi) Updatecandidate(ctx *fiber.Ctx) error {
	var candidate dto.CandidateDtoRequest

	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}

	if err := ctx.BodyParser(&candidate); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "error when parsing data"})
	}

	if err := c.service.UpdateCandidate(id, candidate); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "candidate updated successfuly"})
}

func (c *candidateApi) Deletecandidate(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}

	if err := c.service.DeleteCandidate(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "candidate deleted successfuly"})
}

func NewCandidateApi(app *fiber.App, service CandidateService, middleware middleware.AuthMiddleware) {
	h := candidateApi{
		service:    service,
		middleware: middleware,
	}

	app.Post("candidate/create", h.middleware.RequireToken("admin"), h.Createcandidate)
	app.Get("candidate/getall", h.middleware.RequireToken("admin"), h.GetAllcandidate)
	app.Put("candidate/:id", h.middleware.RequireToken("admin"), h.Updatecandidate)
	app.Delete("candidate/:id", h.middleware.RequireToken("admin"), h.Deletecandidate)

}
