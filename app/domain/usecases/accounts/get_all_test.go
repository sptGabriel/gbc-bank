package accounts

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll(t *testing.T) {
	fakeCpf, _ := vos.NewCPF("70405270062")
	fakeName, _ := vos.NewName("testing create account")
	fakeSecret, _ := vos.NewSecret("123456789")

	fakeAccount := accounts.NewAccount(fakeName, fakeCpf, fakeSecret)

	t.Run("Should get account list", func(t *testing.T) {
		setupUseCaseTest()

		expected := []accounts.Account{fakeAccount}

		mockedRepository.GetAllFunc = func(ctx context.Context) ([]accounts.Account, error) {
			return expected, nil
		}

		accounts, err := mockedUseCase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, expected, accounts)
	})

	t.Run("Should return error when fails to call get Accounts", func(t *testing.T) {
		setupUseCaseTest()

		expected := errors.New("failed")

		mockedRepository.GetAllFunc = func(ctx context.Context) ([]accounts.Account, error) {
			return nil, expected
		}

		accounts, err := mockedUseCase.GetAll(context.Background())

		assert.Equal(t, expected, err)
		assert.Nil(t, accounts)
	})
}