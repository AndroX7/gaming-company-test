package mysql

import (
	"gaming-company-test/models"
	"gaming-company-test/utils/errors"
	"log"

	"gaming-company-test/lib/request_util"
)

func (r *Repository) FindAll(config request_util.PaginationConfig) ([]models.Artist, error) {
	results := []models.Artist{}
	err := r.db.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-all-artists:", err)
		return nil, errors.ErrUnprocessableEntity
	}
	return results, nil
}
