package catering

import (
	"context"
	"errors"

	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/commands"
)

// Service for catering
type Service struct {
	bus       base.Bus
	orders    base.Repository
	readModel *readModel
}

// NewService creates a new catering service
func NewService(bus base.Bus) *Service {
	service := &Service{
		bus:       bus,
		orders:    base.NewRepository(&Order{}, bus),
		readModel: newReadModel(),
	}

	bus.Register(context.Background(), service)

	return service
}

// On event
func (s *Service) On(ctx context.Context, event base.Event) error {
	return s.readModel.On(ctx, event)
}

// Apply command
func (s *Service) Apply(ctx context.Context, command base.Command) error {
	// pre-aggregate-apply-hook
	switch c := command.(type) {
	case commands.OrderFood:
		// check if user has enough money (eventually consistent via projection view)
		// check if user can order for user/tenant in command

		if s.readModel.userBalanceFor(c.User) <= 0 {
			return errors.New("balance is too low")
		}

		// order id is based on date and tenant id: projectAXY-2018-10-09
		// should tenant and user be put into the command here or even earlier?
		// YES, should be in the command! another user could place the order, we don't
		// want to rely on the context!
		orderID := orderIDFrom(c.Place, c.Date)

		if s.readModel.IsOrderFrozen[orderID] {
			return errors.New("order is frozen")
		}

		return s.orders.Apply(ctx, orderID, c)
	default:
		return nil
	}
}
