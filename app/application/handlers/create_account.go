package handlers

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type CreateAccountHandler struct {
	repository repositories.AccountRepository
	hasher     ports.Hasher
}

func NewCreateAccountHandler(r repositories.AccountRepository, h ports.Hasher) *CreateAccountHandler {
	return &CreateAccountHandler{
		repository: r,
		hasher:     h,
	}
}

func (ac *CreateAccountHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.CreateAccountCommand)
	if !ok {
		return nil, app.NewInternalError("invalid command", nil)
	}

	name, err := vos.NewName(command.Name)
	if err != nil {
		return nil, err
	}

	if err := ac.hasher.Hash(&command.Secret); err != nil {
		return nil, err
	}

	secret, err := vos.NewSecret(command.Secret)
	if err != nil {
		return nil, err
	}

	cpf, err := vos.NewCpf(command.Cpf)
	if err != nil {
		return nil, err
	}

	hasAccount, err := ac.repository.DoesAccountExistByCPF(ctx, cpf)
	if err != nil {
		return nil, err
	}
	if hasAccount {
		return nil, err
	}

	account := entities.NewAccount(name, cpf, secret)
	if err = ac.repository.Create(ctx, &account); err != nil {
		return nil, err
	}

	return nil, nil
}
