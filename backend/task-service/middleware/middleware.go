package middleware

import (
	"github.com/erloma/manaflo/backend/task-service/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")

		// Check if the header exists and has the right format
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{"error": "Authentication required"})
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		userID, err := auth.ValidateToken(tokenString)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		// Store user ID in context for later use
		c.Locals("userID", userID)

		// Continue to the next middleware/handler
		return c.Next()
	}
}
