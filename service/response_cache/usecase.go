package response_cache

import "gaming-company-test/models"

type Usecase interface {
	FlushFromArtist(artist *models.Artist)

	FlushAllFromSet(groupSet string)
	FlushGeneralSet(groupSet string)
	FlushGeneralSetWithID(groupSet string, ID uint64)
	FlushCustomSet(groupSet string, customFieldValue string)
}
