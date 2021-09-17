package usecase

import "gaming-company-test/models"

func (u *Usecase) Show(artistID uint64) (*models.Artist, error) {
	artist, err := u.artistRepo.FindByID(artistID)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
