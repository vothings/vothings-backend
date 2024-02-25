package api

import (
	"log"
	"vothings/internal/app/dto"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Register(req dto.UserRegisterReq) error
	Login(req dto.UserLoginReq) (string, error)
}

type AuthApi struct {
	userService UserService
}

func NewAuth(app *fiber.App, userService UserService) {
	h := AuthApi{
		userService: userService,
	}
	app.Post("user/register", h.Register)
	app.Post("user/login", h.Login)

}

func (a *AuthApi) Register(ctx *fiber.Ctx) error {
	var user dto.UserRegisterReq

	if err := ctx.BodyParser(&user); err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{"message": "failed to read data"})
	}

	if err := a.userService.Register(user); err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "register successfuly"})
}

func (a *AuthApi) Login(ctx *fiber.Ctx) error {
	var userRequest dto.UserLoginReq

	if err := ctx.BodyParser(&userRequest); err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{"message": "failed to read data"})
	}

	token, err := a.userService.Login(userRequest)

	if err != nil {
		log.Println(err)
		return ctx.SendStatus(400)
	}

	return ctx.Status(200).JSON(fiber.Map{"token": token})
}
