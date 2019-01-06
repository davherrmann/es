package catering

// FoodOrdered event
type FoodOrdered struct {
	User  string
	Place string
	Date  string
	Food  string
}

// FoodOrderCancelled event
type FoodOrderCancelled struct {
	User  string
	Place string
	Date  string
}

// OrderFrozen event
type OrderFrozen struct {
	Place string
	Date  string
}
