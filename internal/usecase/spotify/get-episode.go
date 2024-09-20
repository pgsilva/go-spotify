package spotify

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/pgsilva/go-spotify/internal/model/spotifyapi"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func GetEpisodes(id string, token string) (spotifyapi.SpotifyShowEpisodeApiResponse, error) {
	slog.Info("Getting episode", "showid", id)

	request, err := prepareGetEpisodesRequest(id, token)
	if err != nil {
		slog.Error("Error preparing request", "err", err)
		return spotifyapi.SpotifyShowEpisodeApiResponse{}, err
	}

	client := config.GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Error calling Spotify API", "err", err)
		return spotifyapi.SpotifyShowEpisodeApiResponse{}, err
	}

	episodes := spotifyapi.SpotifyShowEpisodeApiResponse{}
	if err := json.Unmarshal(response, &episodes); err != nil {
		return spotifyapi.SpotifyShowEpisodeApiResponse{}, err
	}

	return episodes, nil

}

func GetEpisode(showId string, episodeId string, token string) (spotifyapi.SpotifyShowEpisode, error) {
	slog.Info("Getting episode", "showid", showId, "episodeid", episodeId)

	request, err := prepareGetEpisodeRequest(episodeId, token)
	if err != nil {
		slog.Error("Error preparing request", "err", err)
		return spotifyapi.SpotifyShowEpisode{}, err
	}

	client := config.GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Error calling Spotify API", "err", err)
		return spotifyapi.SpotifyShowEpisode{}, err
	}

	episode := spotifyapi.SpotifyShowEpisode{}
	if err := json.Unmarshal(response, &episode); err != nil {
		return spotifyapi.SpotifyShowEpisode{}, err
	}

	return episode, nil
}

func prepareGetEpisodesRequest(id string, token string) (*http.Request, error) {
	slog.Info("Preparing request to call Spotify API")

	apiPath, err := url.Parse(config.SpotifyContentApiUrl + "shows/" + id + "/episodes")
	if err != nil {
		slog.Error("Error parsing URL", "err", err)
		return nil, err
	}

	params := url.Values{}
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

func prepareGetEpisodeRequest(episodeId string, token string) (*http.Request, error) {
	slog.Info("Preparing request to call Spotify API")

	apiPath, err := url.Parse(config.SpotifyContentApiUrl + "episodes/" + episodeId)
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
