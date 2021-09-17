package models

import (
	"gorm.io/gorm"
	"time"
)

type Artist struct {
	ID          uint64          `json:"id"`
	ArtistName  string          `json:"artist_name"`
	AlbumName   string          `json:"album_name"`
	ImageUrl    string          `json:"image_url"`
	ReleaseDate *time.Time      `json:"release_date"`
	Price       float64         `json:"price"`
	SampleUrl   string          `json:"sample_url"`
	CreatedAt   time.Time       `json:"-"`
	UpdatedAt   time.Time       `json:"-"`
	DeletedAt   *gorm.DeletedAt `json:"-"`
}
