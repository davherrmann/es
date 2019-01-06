package catering

import (
	"context"
	"log"

	"github.com/davherrmann/es/api/catering"
	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/service/catering/resolver"
)

// Catering queries
type Catering struct {
	Service Service
}

// Orders resolution
func (r *Catering) Orders(ctx context.Context) []resolver.Order {
	res := r.Service.readModel.Orders

	orders := []resolver.Order{}

	for _, order := range res {
		orders = append(orders, resolver.Order{Order: order})
	}

	return orders
}

// OrderFood mutation
func (r *Catering) OrderFood(ctx context.Context, command catering.DoOrderFood) (bool, error) {
	return r.publish(ctx, command)
}

// CancelFoodOrder mutation
func (r *Catering) CancelFoodOrder(ctx context.Context, command catering.DoCancelFoodOrder) (bool, error) {
	return r.publish(ctx, command)
}

func (r *Catering) publish(ctx context.Context, command base.Command) (bool, error) {
	log.Printf("publishing %#v", command)
	return true, r.Service.Apply(ctx, command)
}
