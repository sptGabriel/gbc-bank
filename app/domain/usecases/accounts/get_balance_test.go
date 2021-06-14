package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	. "github.com/sptGabriel/banking/app/gateway/db/postgres/accounts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	fakeCpf, _ := vos.NewCPF("70405270062")
	fakeName, _ := vos.NewName("testing create account")
	fakeSecret, _ := vos.NewSecret("123456789")

	fakeAccount := accounts.NewAccount(fakeName, fakeCpf, fakeSecret)

	t.Run("Should get account balance", func(t *testing.T) {
		setupUseCaseTest()

		mockedRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*accounts.Account, error) {
			return &fakeAccount, nil
		}

		id := fakeAccount.Id

		account, err := mockedUseCase.GetBalance(context.Background(), id)

		assert.Nil(t, err)
		assert.Equal(t, &fakeAccount, account)
	})

	t.Run("Should return error when fails to call get account balance", func(t *testing.T) {
		setupUseCaseTest()

		mockedRepository.GetByIDFunc = func(ctx context.Context, cpf vos.AccountId) (*accounts.Account, error) {
			return nil, ErrAccountNotFound
		}

		account, err := mockedUseCase.GetBalance(context.Background(), fakeAccount.Id)

		assert.Equal(t, ErrAccountNotFound, err)
		assert.Nil(t, account)
	})
}