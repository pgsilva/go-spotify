package requests

type FavoriteShow struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Uri         string `json:"uri"`
	PosterImage string `json:"poster_image"`
}
