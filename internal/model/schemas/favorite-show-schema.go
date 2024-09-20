package schemas

import (
	"database/sql"

	"gorm.io/gorm"
)

type FavoriteShow struct {
	gorm.Model

	ID          uint           `gorm:"primaryKey;autoIncrement"`
	IdShow      string         `gorm:"not null"`
	Name        string         `gorm:"not null"`
	Description sql.NullString `gorm:"null"`
	Publisher   sql.NullString `gorm:"null"`
	Uri         string         `gorm:"not null"`
	PosterImage string         `gorm:"not null"`
}
