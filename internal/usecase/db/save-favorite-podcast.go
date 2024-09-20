package db

import (
	"database/sql"

	"github.com/pgsilva/go-spotify/internal/model/requests"
	"github.com/pgsilva/go-spotify/internal/model/schemas"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func CreateFavoriteShow(request requests.FavoriteShow) (schemas.FavoriteShow, error) {
	db := config.DB

	//TODO validate request

	favoriteShow := schemas.FavoriteShow{
		IdShow:      request.ID,
		Name:        request.Name,
		Description: sql.NullString{String: request.Description, Valid: request.Description != ""},
		Publisher:   sql.NullString{String: request.Publisher, Valid: request.Publisher != ""},
		Uri:         request.Uri,
		PosterImage: request.PosterImage,
	}

	if err := db.Create(&favoriteShow).Error; err != nil {
		return favoriteShow, err
	}

	return favoriteShow, nil
}
