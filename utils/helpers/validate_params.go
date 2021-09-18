package helpers

import "gaming-company-test/utils/errors"

func ValidateParams(artistID uint64) error {

	if artistID < 1 {
		err := errors.ErrUnprocessableEntity
		err.Message = "invalid params"
		return err
	}
	return nil
}
