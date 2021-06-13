package services

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	accRepo "github.com/sptGabriel/banking/app/gateway/db/postgres/accounts"
	"github.com/sptGabriel/banking/app/gateway/ports"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type authenticate struct {
	accRepo accounts.Repository
	hash    ports.Hash
	cipher  ports.Cipher
}

type Authenticate interface {
	Authenticate(ctx context.Context, cpf vos.CPF, secret vos.Secret) (string, error)
}

func NewAuth(repo accounts.Repository, hash ports.Hash, cipher ports.Cipher) *authenticate {
	return &authenticate{repo, hash, cipher}
}

func (a *authenticate) Authenticate(ctx context.Context, cpf vos.CPF, secret vos.Secret) (string, error) {
	const operation = "Domain.Services.Authenticate"
	account, err := a.accRepo.GetByCPF(ctx, cpf)
	if err != nil {
		if errors.Is(err, accRepo.ErrAccountNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", app.Err(operation, err)
	}

	plainSecret := []byte(secret.String())
	hashedSecret := []byte(account.Secret.String())

	if err := a.hash.Compare(hashedSecret, plainSecret); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := a.cipher.Encrypt(account.Id.String())
	if err := a.hash.Compare(hashedSecret, plainSecret); err != nil {
		return "", app.Err(operation, err)
	}

	return token, nil
}
