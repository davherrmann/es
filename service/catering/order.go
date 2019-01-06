package catering

import (
	"context"
	"errors"
	"fmt"

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
	case command.CancelFoodOrder:
		return []base.Event{event.FoodOrderCancelled{
			User:  c.User,
			Place: c.Place,
			Date:  c.Date,
		}}, nil
	}

	return base.NoEvents()
}

// helpers
func orderIDFrom(place string, date string) string {
	return fmt.Sprintf("%s-%s", place, date)
}
