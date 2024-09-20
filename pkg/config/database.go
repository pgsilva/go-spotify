package config

import (
	"fmt"
	"log/slog"

	"github.com/pgsilva/go-spotify/internal/model/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() error {
	host := PostgreSQLHost
	port := PostgreSQLPort
	user := PostgreSQLUser
	password := PostgreSQLPassword
	dbname := PostgreSQLDatabase

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	slog.Info("Connecting to PostgreSQL database", "dsn", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to the database", "err", err)
		return err
	}

	err = db.AutoMigrate(&schemas.FavoriteDevice{})
	if err != nil {
		slog.Error("Failed to migrate the database", "err", err, "table", "favorite_device")
		return err
	}

	err = db.AutoMigrate(&schemas.FavoriteShow{})
	if err != nil {
		slog.Error("Failed to migrate the database", "err", err, "table", "favorite_show")
		return err
	}

	DB = db
	slog.Info("Database connection established successfully")
	return nil
}
