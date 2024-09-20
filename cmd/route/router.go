package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/internal/handler"
	"github.com/pgsilva/go-spotify/internal/middleware"
)

const (
	ApiV1BasePath   = "/api/v1"
	SpodifyBasePath = "/spodify"
)

func InitRoutes(router *fiber.App) {

	v1 := router.Group(ApiV1BasePath)

	{
		v1.Get("/health", handler.HealthCheck)
	}

	spodifyRoutes := v1.Group(SpodifyBasePath)
	spodifyRoutes.Use([]string{"/player"}, middleware.GenerateApiToken)
	{
		spodifyRoutes.Get("/player/search", handler.SearchShow)
	}
}
