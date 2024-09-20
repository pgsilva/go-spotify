package db

import (
	"log/slog"

	"github.com/pgsilva/go-spotify/internal/model/schemas"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func GetFavoriteShow() ([]schemas.FavoriteShow, error) {
	slog.Info("Getting favorite show in database")
	db := config.DB
	var favoriteShows []schemas.FavoriteShow
	err := db.Find(&favoriteShows).Error
	return favoriteShows, err
}
