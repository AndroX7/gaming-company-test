package mysql

import (
	"gaming-company-test/models"
	"gaming-company-test/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Update(artist *models.Artist, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(artist).Error
	if err != nil {
		log.Println("error-update-artist:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
