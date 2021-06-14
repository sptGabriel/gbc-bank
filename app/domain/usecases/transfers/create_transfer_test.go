package transfers

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	fakeOriginId := vos.NewAccountId()
	fakeDestinationId := vos.NewAccountId()
	fakeAmount := 10000
	// fakeTransfer, _ = transfers.NewTransfer(fakeOriginId, fakeDestinationId, fakeAmount)

	t.Run("Should create an transfers", func(t *testing.T) {
		setupUseCaseTest()

		mockedTransactional.ExecFunc = func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, nil
		}

		transfer, _ := transfers.NewTransfer(fakeOriginId, fakeDestinationId, fakeAmount)
		err := mockedUseCase.CreateTransfer(context.Background(), transfer)

		assert.Nil(t, err)
	})

	t.Run("Should not create transfer, same ids", func(t *testing.T) {
		setupUseCaseTest()

		mockedTransactional.ExecFunc = func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, transfers.ErrSELFTransfer
		}

		transfer, _ := transfers.NewTransfer(fakeOriginId, fakeDestinationId, fakeAmount)
		err := mockedUseCase.CreateTransfer(context.Background(), transfer)

		assert.Equal(t, transfers.ErrSELFTransfer, err)
	})


	t.Run("Should not create transfer, origin account not found", func(t *testing.T) {
		setupUseCaseTest()

		mockedTransactional.ExecFunc = func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, ErrAccountOriginNotFound
		}

		transfer, _ := transfers.NewTransfer(fakeOriginId, fakeDestinationId, fakeAmount)
		err := mockedUseCase.CreateTransfer(context.Background(), transfer)

		assert.Equal(t, ErrAccountOriginNotFound, err)
	})

	t.Run("Should not create transfer, destination account not found", func(t *testing.T) {
		setupUseCaseTest()

		mockedTransactional.ExecFunc = func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, ErrAccountDestinationNotFound
		}

		transfer, _ := transfers.NewTransfer(fakeOriginId, fakeDestinationId, fakeAmount)
		err := mockedUseCase.CreateTransfer(context.Background(), transfer)

		assert.Equal(t, ErrAccountDestinationNotFound, err)
	})
}