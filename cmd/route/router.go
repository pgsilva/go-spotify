package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/internal/handler"
	"github.com/pgsilva/go-spotify/internal/middleware"
)

const (
	ApiV1BasePath   = "/v1/api"
	SpodifyBasePath = "/spodify"
)

func InitRoutes(router *fiber.App) {

	/* /v1/api */
	v1 := router.Group(ApiV1BasePath)

	{
		v1.Get("/health", handler.HealthCheck)
	}

	/* /v1/api/spodify */
	spodifyRoutes := v1.Group(SpodifyBasePath)
	spodifyRoutes.Use([]string{"/podcast"}, middleware.GenerateApiToken)
	{
		spodifyRoutes.Get("/podcast/search", handler.SearchShow)
		spodifyRoutes.Get("/podcast/:id", handler.GetShow)
		spodifyRoutes.Get("/podcast/:id/episode", handler.GetEpisodes)
		spodifyRoutes.Get("/podcast/:id/episode/:episode_id", handler.GetEpisodes)

		// spodifyRoutes.Put("/podcast/:id/play", handler.PlayShow)
		// spodifyRoutes.Put("/podcast/:id/pause", handler.PauseShow)

		// spodifyRoutes.Get("/podcast/me/devices", handler.GetDevices)

	}

	/* /v1/api/spodify/db */
	dbRoutes := spodifyRoutes.Group("/db")
	{
		// dbRoutes.Get("/device", handler.GetFavoriteDevice)
		// dbRoutes.Post("/device", handler.CreateFavoriteDevice)
		// dbRoutes.Delete("/device/:id", handler.DeleteFavoriteDevice)

		dbRoutes.Get("/podcast", handler.GetFavoriteShow)
		dbRoutes.Post("/podcast", handler.CreateFavoriteShow)
		dbRoutes.Delete("/podcast/:id", handler.DeleteFavoriteShow)
	}
}
