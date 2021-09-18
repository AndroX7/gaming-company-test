package events

import (
	"gaming-company-test/models"
)

const (
	EventNameArtistUpdated = "artist_updated"
)

type Client interface {
	DispatchArtistUpdatedEvent(artistM *models.Artist) error
}

type Event struct {
}

func Register() Client {
	return &Event{}
}
