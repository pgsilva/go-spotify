package spotify

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/pgsilva/go-spotify/internal/model/spotifyapi"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func GetShow(id string, token string) (spotifyapi.SpotifyShow, error) {
	slog.Info("Getting show", "id", id)

	request, err := prepareGetShowRequest(id, token)
	if err != nil {
		slog.Error("Error preparing request", "err", err)
		return spotifyapi.SpotifyShow{}, err
	}

	client := config.GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Error calling Spotify API", "err", err)
		return spotifyapi.SpotifyShow{}, err
	}

	show := spotifyapi.SpotifyShow{}
	if err := json.Unmarshal(response, &show); err != nil {
		return spotifyapi.SpotifyShow{}, err
	}

	return show, nil
}

func prepareGetShowRequest(id string, token string) (*http.Request, error) {
	slog.Info("Preparing request to call Spotify API")

	apiPath, err := url.Parse(config.SpotifyContentApiUrl + "shows/" + id)
	if err != nil {
		slog.Error("Error parsing URL", "err", err)
		return nil, err
	}

	params := url.Values{}
	params.Add("market", "BR")
	apiPath.RawQuery = params.Encode()

	slog.Info("Request data", "url", apiPath, "data", params)

	req, err := http.NewRequest("GET", apiPath.String(), nil)
	if err != nil {
		slog.Error("Error creating request", "err", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	return req, nil
}
