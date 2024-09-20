package spotifyapi

type SpotifyShowApiResponse struct {
	Shows struct {
		Href     string        `json:"href"`
		Limit    int           `json:"limit"`
		Next     string        `json:"next"`
		Offset   int           `json:"offset"`
		Previous interface{}   `json:"previous"`
		Total    int           `json:"total"`
		Items    []SpotifyShow `json:"items"`
	} `json:"shows"`
}

type SpotifyShow struct {
	// AvailableMarkets []string      `json:"available_markets"`
	Copyrights      []interface{} `json:"copyrights"`
	Description     string        `json:"description"`
	HTMLDescription string        `json:"html_description"`
	Explicit        bool          `json:"explicit"`
	ExternalUrls    struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	IsExternallyHosted bool     `json:"is_externally_hosted"`
	Languages          []string `json:"languages"`
	MediaType          string   `json:"media_type"`
	Name               string   `json:"name"`
	Publisher          string   `json:"publisher"`
	Type               string   `json:"type"`
	URI                string   `json:"uri"`
	TotalEpisodes      int      `json:"total_episodes"`
}
