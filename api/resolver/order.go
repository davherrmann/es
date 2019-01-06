package resolver

import (
	"context"

	"github.com/davherrmann/es/query"
)

// Order resolver
type Order struct {
	order query.Order
}

// NewOrder resolver
func NewOrder(ctx context.Context, query query.Query, orderID string) (*Order, error) {
	order, err := query.FindOrderByID(ctx, orderID)

	if err != nil {
		return nil, err
	}

	return &Order{
		order: order,
	}, nil
}

// Place resolution
func (o Order) Place() string {
	return o.order.Place
}

// User resolution
func (o Order) User() string {
	return o.order.User
}

// Date resolution
func (o Order) Date() string {
	return o.order.Date
}

// Food resolution
func (o Order) Food() string {
	return o.order.Food
}
