package catering

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/command"
	"github.com/davherrmann/es/event"
)

// Order aggregate
type Order struct {
	OrderCount int
	Orders     map[string]*event.FoodOrdered // user -> event
}

// On rehydrates the aggregate
func (a *Order) On(ctx context.Context, evt base.Event) error {
	switch e := evt.(type) {
	case *event.FoodOrdered:
		a.Orders[e.User] = e
		a.OrderCount++
	}

	return nil
}

// Apply runs strong consistency checks
// Eventual consistency checks should be run in the service (pre-apply-hook)
// Input validation should be done in command.Validate()
func (a *Order) Apply(ctx context.Context, cmd base.Command) ([]base.Event, error) {
	switch c := cmd.(type) {
	case command.OrderFood:
		if a.OrderCount > 10 {
			return nil, errors.New("order limit reached (10), can't add order")
		}

		return []base.Event{event.FoodOrdered{
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
