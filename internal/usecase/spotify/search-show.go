package spotify

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/pgsilva/go-spotify/internal/model/spotifyapi"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func SearchShow(query string, token string) (spotifyapi.SpotifyShowApiResponse, error) {
	slog.Info("Searching for show", "query", query, "token", token)

	request, err := prepareRequest(query, token)
	if err != nil {
		slog.Error("Error preparing request", "err", err)
		return spotifyapi.SpotifyShowApiResponse{}, err
	}

	client := config.GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Error calling Spotify API", "err", err)
		return spotifyapi.SpotifyShowApiResponse{}, err
	}

	show := spotifyapi.SpotifyShowApiResponse{}
	if err := json.Unmarshal(response, &show); err != nil {
		return spotifyapi.SpotifyShowApiResponse{}, err
	}

	return show, nil
}

func prepareRequest(query string, token string) (*http.Request, error) {
	slog.Info("Preparing request to call Spotify API")

	apiPath, err := url.Parse(config.SpotifyContentApiUrl + "search")
	if err != nil {
		slog.Error("Error parsing URL", "err", err)
		return nil, err
	}

	params := url.Values{}
	params.Add("q", query)
	params.Add("type", "show")
	params.Add("market", "BR")
	params.Add("limit", "20")
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
