package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/pgsilva/go-spotify/internal/model/spotifyapi"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func GenerateApiToken(c fiber.Ctx) error {
	/*
		Before call any endpoint Spotify api we need to generate a token to authenticate the request
	*/

	slog.Info("Generating Spotify API token")

	uuid := uuid.New()
	slog.Info("Request ID", "id", uuid)

	request, err := prepareRequest()
	if err != nil {
		slog.Error("Error preparing request", "err", err)
		responseError(c, err)
	}

	client := config.GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Error calling Spotify API", "err", err)
		responseError(c, err)

	}

	slog.Info("Spotify API token generated successfully")
	slog.Info("Storing token in context")

	token := spotifyapi.SpotifyTokenResponse{}

	if err := json.Unmarshal(response, &token); err != nil {
		slog.Error("Error parsing response", "err", err)
		responseError(c, err)

	}

	// Armazena o token no contexto do Fiber
	c.Locals(spotifyapi.SpotifyTokenKey, token.AccessToken)
	return c.Next()
}

func prepareRequest() (*http.Request, error) {
	slog.Info("Preparing request to generate Spotify API token")

	apiPath := config.SpotifyAccountApiUrl + "token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", config.SpotifyClientId)
	data.Set("client_secret", config.SpotifyClientSecret)

	encodedData := data.Encode()

	slog.Info("Request data", "url", apiPath, "data", encodedData)

	req, err := http.NewRequest("POST", apiPath, strings.NewReader(encodedData))
	if err != nil {
		slog.Error("Error creating request", "err", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func responseError(c fiber.Ctx, err error) {
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Error preparing request",
		"error":   err,
	})
}
