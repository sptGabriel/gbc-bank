package mediator

import (
	"context"
	"errors"
	"reflect"
)

var (
	ErrCommandAlreadyRegistered = errors.New("command is already registered on this bus")
	ErrNotRegisteredHandler     = errors.New("handler is not registered on the bus")
)

type Command interface {
}

type Handler interface {
	Execute(ctx context.Context, command Command) (interface{}, error)
}

type Bus struct {
	handlers map[reflect.Type]Handler
}

func NewBus() Bus {
	return Bus{handlers: make(map[reflect.Type]Handler)}
}

func (cb *Bus) RegisterHandler(c Command, ch Handler) error {
	cmdName := reflect.TypeOf(c)
	if _, has := cb.handlers[cmdName]; has {
		return ErrCommandAlreadyRegistered
	}
	cb.handlers[cmdName] = ch
	return nil
}

func (cb Bus) Publish(ctx context.Context,c Command) (interface{}, error) {
	cmdName := reflect.TypeOf(c)
	ch, ok := cb.handlers[cmdName]
	if !ok {
		return nil, ErrNotRegisteredHandler
	}
	return ch.Execute(ctx, c)
}
