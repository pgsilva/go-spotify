package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-spotify/cmd/route"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func main() {
	startupEnvVariables()
	startupDb()
	startupApp()
}

func startupEnvVariables() {
	slog.Info("Starting the environment variables...")
	if err := config.InitEnvVariables(); err != nil {
		return
	}
}

func startupDb() error {
	slog.Info("Starting the database...")

	err := config.InitPostgres()
	if err != nil {
		slog.Error("Failed to start the database", "err", err)
		return err
	}

	slog.Info("Database connection established successfully")

	return nil
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
