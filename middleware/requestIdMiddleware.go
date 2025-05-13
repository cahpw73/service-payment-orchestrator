package middleware

import (
	"service-payment-orchestrator/utils"

	"github.com/gofiber/fiber/v2"
)

func RequestIDMiddleware(c *fiber.Ctx) error {
	requestID := utils.GenerateRequestID(16)
	c.Locals("requestID", requestID)
	return c.Next()
}
