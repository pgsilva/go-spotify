package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/cmd/route"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func main() {
	startupEnvVariables()
	startupApp()
}

func startupEnvVariables() {
	if err := config.InitEnvVariables(); err != nil {
		return
	}
}

func startupApp() {
	slog.Info("Starting the application...")

	app := fiber.New()
	app.Listen(config.Port)

	route.InitRoutes(app)

	slog.Info("Application started successfully")

	if err := app.Listen(":" + config.Port); err != nil {
		slog.Error("Failed to start the application", "err", err)
	}

}
