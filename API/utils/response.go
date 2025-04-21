package utils

import (
	"github.com/gofiber/fiber/v2"
)

// MessageJSON untuk mengirim response JSON dengan status dan pesan.
func MessageJSON(ctx *fiber.Ctx, statusCode int, status string, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}
