package abstractions

import "github.com/iscfgibarra/applabs-data/events"

type EventBus interface {
	PushEvent(*events.Event)
	Push(string, interface{}, events.EventType)
	Pull() events.Events
	PushEvents(events.Events)
}

type EventBusPublisher interface {
	Publish(events.Events)
}
