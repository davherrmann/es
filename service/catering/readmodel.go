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
				Date:  "asdf",
				Food:  "Pommes",
				Place: "X",
				User:  "B",
			},
			query.Order{
				Date:  "weff",
				Food:  "Pommes",
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
		r.Orders = append(r.Orders, query.Order{
			Date:  e.Date.String(),
			Food:  e.Food,
			Place: e.Place,
			User:  e.User,
		})
	}

	return nil
}

func (r *readModel) userBalanceFor(user string) cents {
	return r.UserBalance[user]
}
