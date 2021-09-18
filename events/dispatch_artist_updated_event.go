package events

import (
	"gaming-company-test/events/entity"
	"gaming-company-test/models"
	"log"

	"github.com/gookit/event"
)

func (e *Event) DispatchArtistUpdatedEvent(artistM *models.Artist) error {
	customEvent := &entity.ArtistEventEntity{ArtistM: artistM}
	customEvent.SetName(EventNameArtistUpdated)
	err := event.FireEvent(customEvent)
	if err != nil {
		log.Printf("error while dispatching %s event: %s\n", EventNameArtistUpdated, err)
		return err
	}

	return nil
}
