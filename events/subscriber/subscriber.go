package subscriber

import "gaming-company-test/listener/event_listeners"

type Client interface {
	SubscribedEvents() map[string]interface{}
}

type Subscriber struct {
	eventListener event_listeners.Client
}

func New(
	eventListener event_listeners.Client,
) Client {
	return &Subscriber{
		eventListener: eventListener,
	}
}
