package service

import (
	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/service/catering"
)

// API combines all services
type API struct {
	catering.Query
}

// New API
func New(bus base.Bus) *API {
	return &API{
		catering.Query{Service: catering.NewService(bus)},
	}
}
