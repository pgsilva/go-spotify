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

	return nil
}
