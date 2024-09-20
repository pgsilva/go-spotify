package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

func HealthCheck(c fiber.Ctx) error {
	slog.Info("Health check endpoint called")

	c.Response().Header.Set("Content-Type", "application/json")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Wake up, Neo...",
	})

	return nil
}
