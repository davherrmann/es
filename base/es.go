package base

import (
	"context"
	"log"
)

// Very basic root types:
// - command handler (aggregate root)
// - event handler (projection, side effects, ...)

// EventHandler (aggregates, services, readmodel projections)
type EventHandler interface {
	On(ctx context.Context, event Event) error
}

// Aggregate interface
type Aggregate interface {
	EventHandler
	// Convention over configuration (return events instead of relying on side effects)
	Apply(ctx context.Context, command Command) ([]Event, error)
}

// Command interface
type Command interface{}

// ValidatedCommand marks a validated command
type ValidatedCommand interface {
	Command
	Validate(ctx context.Context) error
}

// EventData interface
type EventData interface{}

// Event should be a struct containing the data, metadata, version, ...?
type Event interface{}

// NoEvents returns empty event list and no error
func NoEvents() ([]Event, error) {
	return []Event{}, nil
}

// Repository rehydrates, applies commands and stores events
type Repository interface {
	Apply(ctx context.Context, aggregateID string, command Command) error
}

// NewRepository creates a standard repository
func NewRepository(prototype Aggregate, bus Bus) Repository {
	return &standardRepository{
		bus:       bus,
		prototype: prototype,
	}
}

type standardRepository struct {
	bus       Bus
	prototype Aggregate
}

func (r *standardRepository) Apply(ctx context.Context, aggregateID string, command Command) error {
	events, err := r.prototype.Apply(ctx, command)

	if err != nil {
		return err
	}

	// only storing in the event store should be done here
	// publishing should be done behind the store
	for _, event := range events {
		r.bus.Publish(ctx, event)
	}

	return nil
}

// Bus for publishing events, registering for events and sending commands
type Bus interface {
	Publish(ctx context.Context, event Event) error
	Register(ctx context.Context, eventHandler EventHandler) error
}

// NewBus creates an in memory bus
func NewBus() Bus {
	return &standardBus{}
}

type standardBus struct {
	eventHandlers []EventHandler
}

func (b *standardBus) Publish(ctx context.Context, event Event) error {
	log.Printf("event: %#v\n", event)

	for _, eventHandler := range b.eventHandlers {
		err := eventHandler.On(ctx, event)
		if err != nil {
			log.Println("error while publishing: " + err.Error())
		}
	}
	return nil
}

func (b *standardBus) Register(ctx context.Context, eventHandler EventHandler) error {
	b.eventHandlers = append(b.eventHandlers, eventHandler)
	return nil
}
