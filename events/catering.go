package events

import "time"

// FoodOrdered event
type FoodOrdered struct {
	User  string
	Place string
	Date  time.Time
	Food  string
}

// OrderFrozen event
type OrderFrozen struct {
	Place string
	Date  time.Time
}
