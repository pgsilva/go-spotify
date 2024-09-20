package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/internal/model/spotifyapi"
	"github.com/pgsilva/go-spotify/internal/usecase/spotify"
)

func SearchShow(c fiber.Ctx) error {
	token, error := RetrieveToken(c)
	if error != nil {
		return error
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

func GetShow(c fiber.Ctx) error {
	token, error := RetrieveToken(c)
	if error != nil {
		return error
	}

	resp, err := spotify.GetShow(c.Params("id"), token)
	if err != nil {
		slog.Error("Failed to get show", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func GetEpisodes(c fiber.Ctx) error {
	token, error := RetrieveToken(c)
	if error != nil {
		return error
	}

	showId := c.Params("id")
	episodeId := c.Params("episode_id")
	if episodeId != "" {
		resp, err := spotify.GetEpisode(showId, episodeId, token)
		if err != nil {
			slog.Error("Failed to get episodes", "err", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(resp)

	} else {
		resp, err := spotify.GetEpisodes(showId, token)
		if err != nil {
			slog.Error("Failed to get episodes", "err", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func RetrieveToken(c fiber.Ctx) (string, error) {
	token, ok := c.Locals(spotifyapi.SpotifyTokenKey).(string)
	if !ok {
		slog.Error("Token not found in context")
		return "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Token not found in context",
		})
	}

	return token, nil
}
