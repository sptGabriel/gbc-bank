package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type createAccountHandler struct {
	repository repositories.AccountRepository
	hasher     ports.HashService
}

func NewCreateAccountHandler(r repositories.AccountRepository, h ports.HashService) *createAccountHandler {
	return &createAccountHandler{
		repository: r,
		hasher:     h,
	}
}

func (ac *createAccountHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	operation := "Handlers.CreateAccount"

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.CreateAccountCommand)
	if !ok {
		return nil, app.Err(operation, errors.New("invalid transfer command"))
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

	hasAccount, err := ac.repository.DoesExistByCPF(ctx, cpf)
	if err != nil {
		return nil, err
	}
	if hasAccount {
		return nil, app.ErrAccountAlreadyExists
	}

	account := entities.NewAccount(name, cpf, secret)
	err = ac.repository.Create(ctx, &account)

	return nil, err
}

func (ac createAccountHandler) Init(bus mediator.Bus) error {
	return bus.RegisterHandler(commands.CreateAccountCommand{}, &ac)
}