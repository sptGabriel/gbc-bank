package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/repositories"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type getAccountBalanceHandler struct {
	repository repositories.AccountRepository
}

func NewGetAccountBalanceHandler(repository repositories.AccountRepository) *getAccountBalanceHandler {
	return &getAccountBalanceHandler{repository}
}

func (h *getAccountBalanceHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	operation := "Handlers.GetAccount"

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.GetAllAccountBalanceCommand)
	if !ok {
		return nil, app.Err(operation, errors.New("invalid transfer command"))
	}

	account, err := h.repository.GetByID(ctx, command.Id)
	if err != nil {
		return nil, err
	}

	return account, err
}

func (h *getAccountBalanceHandler) Init(bus mediator.Bus) error {
	return bus.RegisterHandler(commands.GetAllAccountBalanceCommand{}, h)
}