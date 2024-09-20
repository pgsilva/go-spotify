package schemas

import "gorm.io/gorm"

type FavoriteDevice struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey;autoIncrement"`
	IdDevice string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Type     string `gorm:"not null"`
}
