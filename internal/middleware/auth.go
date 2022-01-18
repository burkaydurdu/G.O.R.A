package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func NewAuthMiddleware() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "123456",
		},
		Realm: "Forbidden",
		Unauthorized: func(c *fiber.Ctx) error {
			return c.
				Status(fiber.StatusForbidden).
				SendString("Forbidden")
		},
		Next: func(c *fiber.Ctx) bool {
			return c.Path() != "/dashboard"
		},
	})
}
