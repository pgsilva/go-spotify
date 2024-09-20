package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/internal/model/requests"
	"github.com/pgsilva/go-spotify/internal/usecase/db"
)

func GetFavoriteShow(c fiber.Ctx) error {
	slog.Info("Getting favorite show")

	resp, err := db.GetFavoriteShow()
	if err != nil {
		slog.Error("Failed to get favorite show", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func CreateFavoriteShow(c fiber.Ctx) error {
	request := requests.FavoriteShow{}

	if err := c.Bind().Body(&request); err != nil {
		slog.Error("Failed to bind request", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	favoriteShow, err := db.CreateFavoriteShow(request)
	if err != nil {
		slog.Error("Failed to create favorite show", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Favorite show created",
		"id":      favoriteShow.ID,
		"idShow":  favoriteShow.IdShow,
	})

}

func DeleteFavoriteShow(c fiber.Ctx) error {
	id := c.Params("id")

	err := db.DeleteFavoriteShow(id)
	if err != nil {
		slog.Error("Failed to delete favorite show", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Favorite show deleted",
	})
}
