package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"vothings/internal/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	Jwt utils.Jwt
}

func (a *AuthMiddleware) RequireToken(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		if token == "" {
			fmt.Println("kena sini")
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "unathorized"})
		}

		tokenString := strings.Replace(token, "Bearer", "", -1)

		if tokenString == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "unathorized"})
		}
		claims, err := a.Jwt.VerivyToken(tokenString)

		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "unathorized"})
		}

		var validRole bool

		for _, role := range roles {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"message": "forbidden ressource"})
		}

		ctx.Next()

		return nil
	}
}

func NewAuthMiddleware(jwt utils.Jwt) *AuthMiddleware {
	return &AuthMiddleware{
		Jwt: jwt,
	}
}
