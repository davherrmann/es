package catering

import (
	"context"

	"github.com/davherrmann/es/api/catering"
	"github.com/davherrmann/es/api/payment"
	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/service/catering/schema"
)

// TODO do and done instead of event and command?

type cents = int

// read model is updated in memory
// one process from time to time saves read model to db, so that when a service starts,
// it can grab a snapshot from the db and fill in the rest
type readModel struct {
	// this map will grow a lot, should be cleared from time to time?
	// on rehydration it should only be filled with orders that are not in the past
	IsOrderFrozen map[string]bool  // order id -> frozen?
	UserBalance   map[string]cents // user id -> balance
	Orders        []schema.Order
}

func newReadModel() *readModel {
	return &readModel{
		IsOrderFrozen: map[string]bool{},
		UserBalance:   map[string]cents{},
		Orders: []schema.Order{
			schema.Order{
				Date:  "10. Nov",
				Food:  "Standard",
				Place: "X",
				User:  "B",
			},
			schema.Order{
				Date:  "11. Nov",
				Food:  "Vegetarisch",
				Place: "X",
				User:  "A",
			}},
	}
}

func (r *readModel) On(ctx context.Context, evt base.Event) error {
	switch e := evt.(type) {
	case payment.MoneyTransferred:
		r.UserBalance[e.From] -= e.Amount
		r.UserBalance[e.To] += e.Amount
	case catering.OrderFrozen:
		r.IsOrderFrozen[orderIDFrom(e.Place, e.Date)] = true
	case catering.FoodOrdered:
		changed := false
		for i, order := range r.Orders {
			if order.Date == e.Date {
				r.Orders[i].Food = e.Food
				changed = true
			}
		}

		if !changed {
			r.Orders = append(r.Orders, schema.Order{
				Date:  e.Date,
				Food:  e.Food,
				Place: e.Place,
				User:  e.User,
			})
		}
	case catering.FoodOrderCancelled:
		for i, order := range r.Orders {
			if order.Date == e.Date {
				r.Orders[i].Food = ""
			}
		}
	}

	return nil
}

func (r *readModel) userBalanceFor(user string) cents {
	return r.UserBalance[user]
}
