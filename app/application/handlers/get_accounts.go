package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/schemas"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type getAccountsHandler struct {
	repository repositories.AccountRepository
}

func NewGetAccountsHandler(repository repositories.AccountRepository) *getAccountsHandler {
	return &getAccountsHandler{repository}
}

func (h *getAccountsHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	operation := "Handlers.GetAccounts"

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if _, ok := cmd.(commands.GetAllAccountsCommand); !ok {
		return nil, app.Err(operation, errors.New("invalid get account command"))
	}

	accounts, err := h.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return h.parseToSchema(accounts), err
}

func (h *getAccountsHandler) parseToSchema(a []entities.Account) []schemas.AccountSchema {
	accounts := make([]schemas.AccountSchema, 0)

	for _, account := range a {
		schemaAccount := schemas.NewAccountSchema(account)
		accounts = append(accounts, schemaAccount)
	}

	return accounts
}

func (h *getAccountsHandler) Init(bus mediator.Bus) error {
	return bus.RegisterHandler(commands.GetAllAccountsCommand{}, h)
}