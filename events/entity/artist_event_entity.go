package entity

import (
	"gaming-company-test/models"

	"github.com/gookit/event"
)

type ArtistEventEntity struct {
	event.BasicEvent
	ArtistM *models.Artist
}
