package usecase

import (
	"gaming-company-test/models"
	"gaming-company-test/utils/helpers"
)

func (u *Usecase) Show(artistID uint64) (*models.Artist, error) {

	// try to avoid sql injection by injection query using single quotes
	err := helpers.ValidateParams(artistID)
	if err != nil {
		return nil, err
	}

	artist, err := u.artistRepo.FindByID(artistID)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
