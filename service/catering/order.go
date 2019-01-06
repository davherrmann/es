package catering

import (
	"context"
	"errors"
	"fmt"

	"github.com/davherrmann/es/api/catering"
	"github.com/davherrmann/es/base"
)

// Order aggregate
type Order struct {
	OrderCount int
	Orders     map[string]catering.FoodOrdered // user -> event
}

// On rehydrates the aggregate
func (a *Order) On(ctx context.Context, evt base.Event) error {
	switch e := evt.(type) {
	case catering.FoodOrdered:
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
	case catering.DoOrderFood:
		if a.OrderCount > 10 {
			return nil, errors.New("order limit reached (10), can't add order")
		}

		return []base.Event{catering.FoodOrdered{
			User:  c.User,
			Place: c.Place,
			Date:  c.Date,
			Food:  c.Food,
		}}, nil
	case catering.DoCancelFoodOrder:
		return []base.Event{catering.FoodOrderCancelled{
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
