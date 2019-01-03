package events

// MoneyTransferred event
type MoneyTransferred struct {
	From   string
	To     string
	Amount int
}
