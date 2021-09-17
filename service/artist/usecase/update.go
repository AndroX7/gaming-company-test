package usecase

import (
	"gaming-company-test/models"
	"gaming-company-test/service/artist/delivery/http/request"
	"log"

	"github.com/jinzhu/copier"
)

func (u *Usecase) Update(request request.ArtistUpdateRequest, artistID uint64) (*models.Artist, error) {
	artistM, err := u.artistRepo.FindByID(artistID)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(artistM, &request)
	if err != nil {
		log.Println("error-for-copy-request-to-product")
		return nil, err
	}

	err = u.artistRepo.Update(artistM, nil)
	if err != nil {
		return nil, err
	}

	return artistM, err
}
