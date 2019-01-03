package projections

// OrderView projection
type OrderView struct {
	id              string
	status          string
	numberOfActions int
}

// can maybe override default behaviour?
// default: for one tenant and one stream?
// override example: for all events

// On rehydrates the projection with the events
// this can be called live for one/more orders by id
// or projected for all orders
func (v *OrderView) On(event string) {
	switch event {
	case "closed":
		v.status = "closed"
		v.numberOfActions++
	}
}
