package middlewares

import (
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(secret string) fiber.Handler {
	return jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{
			Key: []byte(secret),
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		},
	})
}
