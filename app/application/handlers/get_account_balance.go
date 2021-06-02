package handlers

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type GetAccountBalanceHandler struct {
	repository repositories.AccountRepository
}

func NewGetAccountBalanceHandler(repository repositories.AccountRepository) *GetAccountBalanceHandler {
	return &GetAccountBalanceHandler{repository}
}

func (ac *GetAccountBalanceHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.GetAllAccountBalanceCommand)
	if !ok {
		return nil, app.NewInternalError("invalid command", nil)
	}

	account, err := ac.repository.GetByID(ctx, command.Id)

	return account, err
}
