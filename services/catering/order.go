package catering

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/commands"
	"github.com/davherrmann/es/events"
)

// Order aggregate
type Order struct {
	OrderCount int
	Orders     map[string]*events.FoodOrdered // user -> event
}

// On rehydrates the aggregate
func (a *Order) On(ctx context.Context, event base.Event) error {
	switch e := event.(type) {
	case *events.FoodOrdered:
		a.Orders[e.User] = e
		a.OrderCount++
	}

	return nil
}

// Apply runs strong consistency checks
// Eventual consistency checks should be run in the service (pre-apply-hook)
// Input validation should be done in command.Validate()
func (a *Order) Apply(ctx context.Context, command base.Command) ([]base.Event, error) {
	switch c := command.(type) {
	case commands.OrderFood:
		if a.OrderCount > 10 {
			return nil, errors.New("order limit reached (10), can't add order")
		}

		return []base.Event{events.FoodOrdered{
			User:  c.User,
			Place: c.Place,
			Date:  c.Date,
			Food:  c.Food,
		}}, nil
	}

	return base.NoEvents()
}

// helpers
func orderIDFrom(place string, date time.Time) string {
	return fmt.Sprintf("%s-%02d-%02d-%02d", place, date.Year(), date.Month(), date.Day())
}
