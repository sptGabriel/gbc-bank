package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/schemas"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/tests/fakedata"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)



func TestGetAccounts(t *testing.T) {
	repository := &repositories.AccountRepositoryMock{}
	acc := fakedata.FakeAccount()
	handler := &getAccountsHandler{
		repository: repository,
	}
	cmd := commands.GetAllAccountsCommand{}

	t.Run("Should get account list", func(t *testing.T) {
		expected := []entities.Account{*acc}

		repository.GetAllFunc = func(ctx context.Context) ([]entities.Account, error) {
			return expected, nil
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, err)
		assert.Equal(t, reflect.TypeOf([]schemas.AccountSchema{}), reflect.TypeOf(res))
	})

	t.Run("Should return an error if handler.execute is called with invalid command", func(t *testing.T) {
		badCmd := struct{}{}

		expectedError := app.Err("Handlers.GetAccounts", errors.New("invalid get accounts command"))

		res, err := handler.Execute(context.Background(), badCmd)
		assert.Nil(t, res)
		assert.Equal(t, reflect.TypeOf(&app.DomainError{}), reflect.TypeOf(err))
		assert.Equal(t, expectedError, err)
	})

	t.Run("Should empty account list", func(t *testing.T) {
		expected := make([]entities.Account, 0)

		repository.GetAllFunc = func(ctx context.Context) ([]entities.Account, error) {
			return expected, nil
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, err)
		assert.Equal(t, []schemas.AccountSchema{}, res)
	})

	t.Run("Should return error when fails to call get Accounts", func(t *testing.T) {
		expected := errors.New("failed")

		repository.GetAllFunc = func(ctx context.Context) ([]entities.Account, error) {
			return nil, expected
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, expected, err)
	})
}
