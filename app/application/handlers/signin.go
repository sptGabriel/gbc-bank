package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/application/schemas"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type signInHandler struct {
	accountRepo   repositories.AccountRepository
	cipherService ports.CipherService
	hashService   ports.HashService
}

func NewSignInHandler(r repositories.AccountRepository, c ports.CipherService, h ports.HashService) *signInHandler {
	return &signInHandler{accountRepo: r, cipherService: c, hashService: h}
}

func (h *signInHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	operation := "Handlers.Signin"

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.SignInCommand)
	if !ok {
		return nil, app.Err(operation, errors.New("invalid signin command"))
	}

	cpf, err := vos.NewCpf(command.Cpf)
	if err != nil {
		return nil, err
	}

	account, err := h.accountRepo.GetByCPF(context.Background(), cpf)
	if err != nil {
		return nil, err
	}

	plainSecret := []byte(command.Secret)
	hashedSecret := []byte(account.Secret.String())

	if err := h.hashService.Compare(hashedSecret, plainSecret); err != nil {
		return nil, err
	}

	token, err := h.cipherService.Encrypt(account.Id.String())
	if err != nil {
		return nil, err
	}
	return schemas.NewTokenSchema(token), nil
}

func (h *signInHandler) Init(bus mediator.Bus) error {
	return bus.RegisterHandler(commands.SignInCommand{}, h)
}
