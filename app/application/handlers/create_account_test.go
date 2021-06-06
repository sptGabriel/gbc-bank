package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/tests/fakedata"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)



func TestCreateAccount(t *testing.T) {
	repository := &repositories.AccountRepositoryMock{}
	hasher := &ports.HashServiceMock{}
	handler := &createAccountHandler{
		hasher: hasher,
		repository: repository,
	}

	t.Run("Should create an account", func(t *testing.T) {
		cmd := fakedata.FakeCreateAccountCMD()

		hasher.HashFunc = func(secret *string) error {
			return nil
		}

		repository.CreateFunc = func(ctx context.Context, en *entities.Account) error {
			return nil
		}
		repository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return false, nil
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	t.Run("Should return an error if handler.execute is called with invalid command", func(t *testing.T) {
		badCmd := struct{}{}

		expectedError := app.Err("Handlers.CreateAccount", errors.New("invalid create account command"))

		res, err := handler.Execute(context.Background(), badCmd)
		assert.Nil(t, res)
		assert.Equal(t, reflect.TypeOf(&app.DomainError{}), reflect.TypeOf(err))
		assert.Equal(t, expectedError, err)
	})
	t.Run("Should return an already exists error", func(t *testing.T) {
		cmd := fakedata.FakeCreateAccountCMD()

		repository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return true, nil
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, app.ErrAccountAlreadyExists, err)
	})
	t.Run("Should return an error when failed to hash password", func(t *testing.T) {
		cmd := fakedata.FakeCreateAccountCMD()
		expectedErr := errors.New("failed to hash")

		hasher.HashFunc = func(secret *string) error {
			return expectedErr
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, expectedErr, err)
	})
}
