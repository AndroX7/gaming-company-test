package subscriber

import (
	"gaming-company-test/events"

	"github.com/gookit/event"
)

func (s *Subscriber) SubscribedEvents() map[string]interface{} {
	return map[string]interface{}{
		events.EventNameArtistUpdated: event.ListenerFunc(s.eventListener.ArtistUpdatedEventListener),
	}
}
