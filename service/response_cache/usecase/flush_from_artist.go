package usecase

import (
	"gaming-company-test/app/api/middleware"
	"gaming-company-test/models"
)

func (u *Usecase) FlushFromArtist(artistM *models.Artist) {
	groupSet := middleware.RedisResponseArtistSet

	u.FlushGeneralSet(groupSet)
	u.FlushCustomSet(groupSet, artistM.ArtistName)
	u.FlushGeneralSetWithID(groupSet, artistM.ID)
}
