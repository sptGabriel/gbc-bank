package ports

import "github.com/sptGabriel/banking/app/infrastructure/mediator"

type Handler interface {
	mediator.Handler
	Init(bus mediator.Bus) error
}