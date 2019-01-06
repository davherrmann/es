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
