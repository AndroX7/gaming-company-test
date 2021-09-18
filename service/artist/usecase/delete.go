package usecase

import "gaming-company-test/utils/helpers"

func (u *Usecase) Delete(artistID uint64) error {

	// try to avoid sql injection by injection query using single quotes
	err := helpers.ValidateParams(artistID)
	if err != nil {
		return err
	}

	artistM, err := u.artistRepo.FindByID(artistID)
	if err != nil {
		return err
	}

	err = u.artistRepo.Delete(artistM, nil)

	if err != nil {
		return err
	}

	return nil
}
