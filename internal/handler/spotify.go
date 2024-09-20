package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/internal/model/spotifyapi"
	"github.com/pgsilva/go-spotify/internal/usecase/spotify"
)

func SearchShow(c fiber.Ctx) error {
	token, ok := c.Locals(spotifyapi.SpotifyTokenKey).(string)
	if !ok {
		slog.Error("Token not found in context")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Token not found in context",
		})
	}

	resp, err := spotify.SearchShow(c.Query("q"), token)
	if err != nil {
		slog.Error("Failed to search", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
