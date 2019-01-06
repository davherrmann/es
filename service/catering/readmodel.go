package catering

import (
	"context"

	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/event"
	"github.com/davherrmann/es/query"
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
	Orders        []query.Order
}

func newReadModel() *readModel {
	return &readModel{
		IsOrderFrozen: map[string]bool{},
		UserBalance:   map[string]cents{},
		Orders: []query.Order{
			query.Order{
				Date:  "10. Nov",
				Food:  "Standard",
				Place: "X",
				User:  "B",
			},
			query.Order{
				Date:  "11. Nov",
				Food:  "Vegetarisch",
				Place: "X",
				User:  "A",
			}},
	}
}

func (r *readModel) On(ctx context.Context, evt base.Event) error {
	switch e := evt.(type) {
	case event.MoneyTransferred:
		r.UserBalance[e.From] -= e.Amount
		r.UserBalance[e.To] += e.Amount
	case event.OrderFrozen:
		r.IsOrderFrozen[orderIDFrom(e.Place, e.Date)] = true
	case event.FoodOrdered:
		changed := false
		for i, order := range r.Orders {
			if order.Date == e.Date {
				r.Orders[i].Food = e.Food
				changed = true
			}
		}

		if !changed {
			r.Orders = append(r.Orders, query.Order{
				Date:  e.Date,
				Food:  e.Food,
				Place: e.Place,
				User:  e.User,
			})
		}
	case event.FoodOrderCancelled:
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
