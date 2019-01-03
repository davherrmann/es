package catering

import (
	"context"

	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/events"
)

type cents = int

// read model is updated in memory
// one process from time to time saves read model to db, so that when a service starts,
// it can grab a snapshot from the db and fill in the rest
type readModel struct {
	// this map will grow a lot, should be cleared from time to time?
	// on rehydration it should only be filled with orders that are not in the past
	IsOrderFrozen map[string]bool  // order id -> frozen?
	UserBalance   map[string]cents // user id -> balance
}

func newReadModel() *readModel {
	return &readModel{
		IsOrderFrozen: map[string]bool{},
		UserBalance:   map[string]cents{},
	}
}

func (r *readModel) On(ctx context.Context, event base.Event) error {
	switch e := event.(type) {
	case events.MoneyTransferred:
		r.UserBalance[e.From] -= e.Amount
		r.UserBalance[e.To] += e.Amount
	case events.OrderFrozen:
		r.IsOrderFrozen[orderIDFrom(e.Place, e.Date)] = true
	}

	return nil
}

func (r *readModel) userBalanceFor(user string) cents {
	return r.UserBalance[user]
}
