package handlers

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/schemas"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type GetAccountsHandler struct {
	repository repositories.AccountRepository
}

func NewGetAccountsHandler(repository repositories.AccountRepository) *GetAccountsHandler {
	return &GetAccountsHandler{repository}
}

func (ac *GetAccountsHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if _, ok := cmd.(commands.GetAllAccountsCommand); !ok {
		return nil, app.NewInternalError("invalid command", nil)
	}

	accounts, err := ac.repository.GetAll(ctx)

	return parseToSchema(accounts), err
}

func parseToSchema(a []entities.Account) []schemas.AccountSchema {
	accounts := make([]schemas.AccountSchema, 0)

	for _, account := range a {
		schemaAccount := schemas.NewAccountSchema(account)
		accounts = append(accounts, schemaAccount)
	}

	return accounts
}
