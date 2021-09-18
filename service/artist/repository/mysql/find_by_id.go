package mysql

import (
	"gaming-company-test/models"
	"gaming-company-test/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) FindByID(artistID uint64) (*models.Artist, error) {
	model := models.Artist{}
	err := r.db.
		Where("id = ?", artistID).
		First(&model).Error

	if err == gorm.ErrRecordNotFound || &model == nil {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-artist-by-id:", err)
		return nil, errors.ErrUnprocessableEntity
	}

	return &model, nil
}
