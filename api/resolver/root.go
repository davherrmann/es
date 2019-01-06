package resolver

import (
	"context"
	"log"

	"github.com/davherrmann/es/query"
)

// Root resolver
type Root struct {
	Query query.Query
}

// Hello resolution
func (*Root) Hello() string {
	return "Hello World!!!"
}

// Orders resolution
func (r *Root) Orders(ctx context.Context) []Order {
	res, err := r.Query.FindOrders(ctx)

	orders := []Order{}

	for _, order := range res {
		orders = append(orders, Order{order})
	}

	if err != nil {
		log.Println(err)
	}
	return orders
}

// OrderFood mutation
func (r *Root) OrderFood(ctx context.Context, args *struct {
	Date string
	Food string
}) (bool, error) {
	return true, r.Query.OrderFood(ctx, args.Date, args.Food)
}

// CancelFoodOrder mutation
func (r *Root) CancelFoodOrder(ctx context.Context, args *struct {
	Date string
}) (bool, error) {
	return true, r.Query.CancelFoodOrder(ctx, args.Date)
}
