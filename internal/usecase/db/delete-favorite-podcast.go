package db

import (
	"github.com/pgsilva/go-spotify/internal/model/schemas"
	"github.com/pgsilva/go-spotify/pkg/config"
)

func DeleteFavoriteShow(id string) error {
	db := config.DB

	if err :=
		db.Where("id_show = ?", id).Delete(&schemas.FavoriteShow{}).Error; err != nil {
		return err
	}

	return nil
}
