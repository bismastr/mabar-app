package db

import (
	"fmt"
	"os"

	"github.com/bismastr/discord-bot/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Client *gorm.DB
}

func NewDatabase() (*Db, error) {
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv(("DB_HOST"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, "5432")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Game{}, &model.Session{})

	return &Db{Client: db}, nil
}
