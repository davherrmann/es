package graphql

import (
	"github.com/davherrmann/es/service/catering"
	"github.com/davherrmann/es/service/payment"
)

// Root resolver
type Root struct {
	catering.Catering
	payment.Payment
}

// Hello resolution
func (*Root) Hello() string {
	return "Hello World!!!"
}
