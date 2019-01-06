package query

import (
	"context"
)

// Order view
type Order struct {
	Place string
	Date  string
	User  string
	Food  string
}

// Query interface
type Query interface {
	FindOrderByID(ctx context.Context, orderID string) (Order, error)
	FindOrders(ctx context.Context) ([]Order, error)
	OrderFood(ctx context.Context, date string, food string) error
	CancelFoodOrder(ctx context.Context, date string) error
}
