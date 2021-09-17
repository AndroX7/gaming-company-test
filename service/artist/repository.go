package artist

import (
	"gaming-company-test/lib/request_util"
	"gaming-company-test/models"

	"gorm.io/gorm"
)

type Repository interface {
	Delete(artist *models.Artist, tx *gorm.DB) error
	FindByID(artistID uint64) (*models.Artist, error)
	FindAll(config request_util.PaginationConfig) ([]models.Artist, error)
	// FindByAlbumName(albumName string) (*models.Artist, error)
	// FindByArtistName(artistName string) (*models.Artist, error)
	Insert(artist *models.Artist, tx *gorm.DB) error
	Update(artist *models.Artist, tx *gorm.DB) error
	Count(config request_util.PaginationConfig) (int64, error)
}
