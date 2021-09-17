package usecase

func (u *Usecase) Delete(artistID uint64) error {
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
