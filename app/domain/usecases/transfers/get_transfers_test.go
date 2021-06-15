package transfers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTransfers(t *testing.T) {
	fakeOriginId := vos.NewAccountId()
	fakeDestinationId := vos.NewAccountId()
	fakeAmount := 10000
	fakeTransfer, _ := transfers.NewTransfer(fakeOriginId, fakeDestinationId, fakeAmount)

	t.Run("Should get account transfers list", func(t *testing.T) {
		setupUseCaseTest()

		expected := []transfers.Transfer{fakeTransfer}

		mockedTransferRepository.GetAllFunc = func(ctx context.Context, accountId vos.AccountId) ([]transfers.Transfer, error) {
			return expected, nil
		}

		transfers, err := mockedUseCase.GetTransfers(context.Background(), fakeOriginId)

		assert.Nil(t, err)
		assert.Equal(t, expected, transfers)
	})

	t.Run("Should return error when fails to call get TRANSFERS", func(t *testing.T) {
		setupUseCaseTest()

		expected := errors.New("failed")

		mockedTransferRepository.GetAllFunc = func(ctx context.Context, accountId vos.AccountId) ([]transfers.Transfer, error) {
			return nil, expected
		}

		transfers, err := mockedUseCase.GetTransfers(context.Background(), fakeOriginId)

		assert.Equal(t, expected, err)
		assert.Nil(t, transfers)
	})
}
