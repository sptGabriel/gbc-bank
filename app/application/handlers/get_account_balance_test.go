package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
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



func TestGetAccountBalance(t *testing.T) {
	repository := &repositories.AccountRepositoryMock{}
	expectedAccount := fakedata.FakeAccount()
	handler := &getAccountBalanceHandler{
		repository: repository,
	}
	id := fakedata.FakeAccountID()
	cmd := commands.NewGetBalanceCommand(id)

	t.Run("Should get an account instance", func(t *testing.T) {
		repository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return expectedAccount, nil
		}

		res, err := handler.Execute(context.Background(), cmd)
		response, _ := res.(schemas.AccountBalance)

		assert.Nil(t, err)
		assert.Equal(t, reflect.TypeOf(schemas.AccountBalance{}), reflect.TypeOf(res))
		if err != nil {
			assert.Equal(t, expectedAccount.Id, response.Id)
			assert.Equal(t, expectedAccount.Balance, response.Balance)
		}
	})

	t.Run("Should return an error if handler.execute is called with invalid command", func(t *testing.T) {
		badCmd := struct{}{}

		expectedError := app.Err("Handlers.GetAccountBalance", errors.New("invalid get account balance command"))

		res, err := handler.Execute(context.Background(), badCmd)
		assert.Nil(t, res)
		assert.Equal(t, reflect.TypeOf(&app.DomainError{}), reflect.TypeOf(err))
		assert.Equal(t, expectedError, err)
	})

	t.Run("Should returns account not found", func(t *testing.T) {
		repository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*entities.Account, error) {
			return nil, app.ErrAccountNotFound
		}

		res, err := handler.Execute(context.Background(), cmd)

		assert.Nil(t, res)
		assert.Equal(t, app.ErrAccountNotFound, err)
	})
}
