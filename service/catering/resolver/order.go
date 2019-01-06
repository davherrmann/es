package resolver

import (
	"github.com/davherrmann/es/service/catering/schema"
)

// Order resolver
type Order struct {
	// TODO don't export, use NewOrder instead?
	Order schema.Order
}

// Place resolution
func (o Order) Place() string {
	return o.Order.Place
}

// User resolution
func (o Order) User() string {
	return o.Order.User
}

// Date resolution
func (o Order) Date() string {
	return o.Order.Date
}

// Food resolution
func (o Order) Food() string {
	return o.Order.Food
}
