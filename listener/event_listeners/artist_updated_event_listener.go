package event_listeners

import (
	"gaming-company-test/app/api/middleware"
	"gaming-company-test/events/entity"

	"github.com/gookit/event"
)

func (l *Listener) ArtistUpdatedEventListener(e event.Event) error {
	artistM := e.(*entity.ArtistEventEntity).ArtistM

	l.responseCacheUsecase.FlushFromArtist(artistM)
	l.responseCacheUsecase.FlushGeneralSet(middleware.RedisResponseArtistSet)

	return nil
}
