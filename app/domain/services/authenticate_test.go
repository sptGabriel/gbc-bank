package services

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	accRepo "github.com/sptGabriel/banking/app/gateway/db/postgres/accounts"
	"github.com/sptGabriel/banking/app/gateway/ports"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	fakeCpf, _ := vos.NewCPF("70405270062")
	fakeName, _ := vos.NewName("testing create account")
	fakeSecret, _ := vos.NewSecret("123456789")

	fakeAccount := accounts.NewAccount(fakeName, fakeCpf, fakeSecret)

	mockedRepository := &accounts.RepositoryMock{}
	mockedHasher := &ports.HashMock{}
	mockedCipher := &ports.CipherMock{}

	mockedAuth := authenticate{
		mockedRepository,
		mockedHasher,
		mockedCipher,
	}

	t.Run("Should authenticate successfully", func(t *testing.T) {
		expected := "stone"

		mockedHasher.CompareFunc = func(hashedPassword []byte, password []byte) error {
			return nil
		}
		mockedRepository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*accounts.Account, error) {
			return &fakeAccount, nil
		}
		mockedCipher.EncryptFunc = func(id string) (string, error) {
			return expected, nil
		}

		token, err := mockedAuth.Authenticate(context.Background(), fakeAccount.CPF, fakeAccount.Secret)

		assert.Equal(t, expected, token)
		assert.Nil(t, err)
	})

	t.Run("Should not authenticate, with ErrInvalidCredentials", func(t *testing.T) {
		mockedRepository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*accounts.Account, error) {
			return nil, accRepo.ErrAccountNotFound
		}

		token, err := mockedAuth.Authenticate(context.Background(), fakeAccount.CPF, fakeAccount.Secret)

		assert.Equal(t, ErrInvalidCredentials, err)
		assert.Equal(t, "", token)
	})

	t.Run("Should not authenticate, with ErrInvalidCredentials", func(t *testing.T) {
		mockedHasher.CompareFunc = func(hashedPassword []byte, password []byte) error {
			return ErrInvalidCredentials
		}
		mockedRepository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*accounts.Account, error) {
			return &fakeAccount, nil
		}

		token, err := mockedAuth.Authenticate(context.Background(), fakeAccount.CPF, fakeAccount.Secret)

		assert.Equal(t, ErrInvalidCredentials, err)
		assert.Equal(t, "", token)
	})
}
