package event_listeners

import (
	"gaming-company-test/service/response_cache"

	"github.com/gookit/event"
)

type Client interface {
	ArtistUpdatedEventListener(e event.Event) error
}

type Listener struct {
	responseCacheUsecase response_cache.Usecase
}

func New(
	responseCacheUsecase response_cache.Usecase,
) Client {
	return &Listener{
		responseCacheUsecase: responseCacheUsecase,
	}
}
