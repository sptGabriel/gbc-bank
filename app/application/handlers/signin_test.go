package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/application/schemas"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/tests/fakedata"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)



func TestSignIn(t *testing.T) {
	repository := &repositories.AccountRepositoryMock{}
	hasher := &ports.HashServiceMock{}
	cipher := &ports.CipherServiceMock{}
	acc := fakedata.FakeAccount()
	handler := &signInHandler{
		accountRepo: repository,
		hashService: hasher,
		cipherService: cipher,
	}
	cmd := commands.SignInCommand{
		Secret: "12345678",
		Cpf: "10757060099",
	}

	t.Run("Should returns token", func(t *testing.T) {
		tokenExpected := "STONE"

		repository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*entities.Account, error) {
			return acc, nil
		}

		hasher.CompareFunc = func(hashedPassword []byte, password []byte) error {
			return nil
		}

		cipher.EncryptFunc = func(id string) (string, error) {
			return tokenExpected, nil
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, err)
		assert.Equal(t, reflect.TypeOf(schemas.TokenSchema{}), reflect.TypeOf(res))
		assert.Equal(t, schemas.NewTokenSchema(tokenExpected), res)
	})

	t.Run("Should return account not found error", func(t *testing.T) {
		repository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*entities.Account, error) {
			return nil, app.ErrAccountNotFound
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, app.ErrAccountNotFound, err)
	})

	t.Run("Should return error when compare hash", func(t *testing.T) {
		expectedErr := errors.New("failed to compare hash")

		repository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*entities.Account, error) {
			return acc, nil
		}

		hasher.CompareFunc = func(hashedPassword []byte, password []byte) error {
			return expectedErr
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Should return error in cipher.Encrypt", func(t *testing.T) {
		errExpected := errors.New("failed to encrypt")

		repository.GetByCPFFunc = func(ctx context.Context, cpf vos.CPF) (*entities.Account, error) {
			return acc, nil
		}

		hasher.CompareFunc = func(hashedPassword []byte, password []byte) error {
			return nil
		}

		cipher.EncryptFunc = func(id string) (string, error) {
			return "", errExpected
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, errExpected, err)
	})

	t.Run("Should return an error if handler.execute is called with invalid command", func(t *testing.T) {
		badCmd := struct{}{}

		expectedError := app.Err("Handlers.Signin", errors.New("invalid signin command"))

		res, err := handler.Execute(context.Background(), badCmd)
		assert.Nil(t, res)
		assert.Equal(t, reflect.TypeOf(&app.DomainError{}), reflect.TypeOf(err))
		assert.Equal(t, expectedError, err)
	})
}
