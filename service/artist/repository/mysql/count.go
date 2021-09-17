package mysql

import (
	"gaming-company-test/lib/request_util"
	"gaming-company-test/models"
	"gaming-company-test/utils/errors"
	"log"
)

func (r *Repository) Count(config request_util.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.Artist{}).
		Scopes(config.Scopes()...).
		Count(&count).Error

	if err != nil {
		log.Println("error-count-artist:", err)
		return 0, errors.ErrUnprocessableEntity
	}
	return count, nil
}
