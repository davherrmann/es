package catering

import (
	"context"

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
