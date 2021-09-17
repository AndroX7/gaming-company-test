package usecase

import (
	"gaming-company-test/models"
	"gaming-company-test/service/artist/delivery/http/request"
)

func (u *Usecase) Create(request request.ArtistCreateRequest) (*models.Artist, error) {
	artistM := &models.Artist{
		ArtistName:  request.ArtistName,
		AlbumName:   request.AlbumName,
		ImageUrl:    request.ImageUrl,
		ReleaseDate: request.ReleaseDate,
		Price:       request.Price,
		SampleUrl:   request.SampleUrl,
	}
	err := u.artistRepo.Insert(artistM, nil)
	if err != nil {
		return nil, err
	}
	return artistM, nil
}
