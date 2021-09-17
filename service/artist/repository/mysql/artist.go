package mysql

import (
	"gaming-company-test/models"
	"gaming-company-test/service/artist"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Model struct {
	productM models.Artist
}

func New(
	db *gorm.DB,
) artist.Repository {
	return &Repository{
		db: db,
	}
}
