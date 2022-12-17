package events

import (
	guuid "github.com/google/uuid"
	"time"
)

type EventType string

const (
	ERROR EventType = "error"
	INFO  EventType = "info"
	DEBUG EventType = "debug"
)

type Event struct {
	ID          string
	Description string
	EventType   EventType
	Data        interface{}
	Register    time.Time
}

func New(description string, data interface{}, eventType EventType) *Event {
	return &Event{
		ID:          guuid.NewString(),
		Description: description,
		Data:        data,
		EventType:   eventType,
		Register:    time.Now(),
	}
}

type Events []*Event

type DomainEventBus struct {
	events Events
}

func NewDomainEventBus() *DomainEventBus {
	return &DomainEventBus{
		events: make(Events, 0),
	}
}

func (ds *DomainEventBus) PushEvent(event *Event) {
	ds.events = append(ds.events, event)
}

func (ds *DomainEventBus) PushEvents(evs Events) {
	for _, event := range evs {
		ds.events = append(ds.events, event)
	}
}

func (ds *DomainEventBus) Push(description string, data interface{}, eventType EventType) {
	event := New(description, data, eventType)
	ds.events = append(ds.events, event)
}

func (ds *DomainEventBus) Pull() Events {
	aux := make(Events, 0)

	for _, event := range ds.events {
		aux = append(aux, event)
	}

	ds.events = make(Events, 0)
	return aux
}
