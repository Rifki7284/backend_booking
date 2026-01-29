package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.Locals("user").(*jwt.Token)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid claims",
			})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "role not found",
			})
		}

		for _, r := range roles {
			if role == r {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden",
		})
	}
}
