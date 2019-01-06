package catering

import (
	"context"

	"github.com/davherrmann/es/command"
	"github.com/davherrmann/es/query"
)

// Query for catering
type Query struct {
	Service *Service
}

// FindOrderByID query
func (*Query) FindOrderByID(ctx context.Context, orderID string) (query.Order, error) {
	return query.Order{
		Place: "X",
		Date:  "2019-10-04",
		User:  "B",
		Food:  orderID,
	}, nil
}

// FindOrders query
func (q *Query) FindOrders(ctx context.Context) ([]query.Order, error) {
	return q.Service.readModel.Orders, nil
}

// OrderFood (TODO via command bus)
func (q *Query) OrderFood(ctx context.Context, date string, food string) error {
	return q.Service.Apply(ctx, command.OrderFood{
		Date:  date,
		Food:  food,
		Place: "X",
		User:  "B",
	})
}

// CancelFoodOrder command
func (q *Query) CancelFoodOrder(ctx context.Context, date string) error {
	return q.Service.Apply(ctx, command.CancelFoodOrder{
		Date:  date,
		Place: "X",
		User:  "B",
	})
}
