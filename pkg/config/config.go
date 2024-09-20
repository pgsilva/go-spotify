package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port                 string
	SpotifyClientId      string
	SpotifyClientSecret  string
	SpotifyContentApiUrl string
	SpotifyAccountApiUrl string
	PostgreSQLHost       string
	PostgreSQLPort       string
	PostgreSQLUser       string
	PostgreSQLPassword   string
	PostgreSQLDatabase   string
)

func InitEnvVariables() error {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file", "err", err)
	}

	Port = os.Getenv("PORT")
	SpotifyClientId = os.Getenv("SPOTIFY_CLIENT_ID")
	SpotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	SpotifyContentApiUrl = os.Getenv("SPOTIFY_CONTENT_API_URL")
	SpotifyAccountApiUrl = os.Getenv("SPOTIFY_ACCOUNT_API_URL")
	PostgreSQLHost = os.Getenv("POSTGRES_HOST")
	PostgreSQLPort = os.Getenv("POSTGRES_PORT")
	PostgreSQLUser = os.Getenv("POSTGRES_USER")
	PostgreSQLPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgreSQLDatabase = os.Getenv("POSTGRES_DB")

	return nil
}
