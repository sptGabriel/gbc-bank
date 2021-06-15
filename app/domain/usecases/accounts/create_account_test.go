package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	fakeCpf, _ := vos.NewCPF("70405270062")
	fakeName, _ := vos.NewName("testing create account")
	fakeSecret, _ := vos.NewSecret("123456789")

	t.Run("Should create an account", func(t *testing.T) {
		setupUseCaseTest()

		mockedHasher.HashFunc = func(secret *string) error {
			return nil
		}
		mockedRepository.CreateFunc = func(ctx context.Context, account *accounts.Account) error {
			return nil
		}
		mockedRepository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return false, nil
		}

		acc := accounts.NewAccount(fakeName, fakeCpf, fakeSecret)
		err := mockedUseCase.CreateAccount(context.Background(), acc)

		assert.Nil(t, err)
	})

	t.Run("Should not create account, account already exists", func(t *testing.T) {
		setupUseCaseTest()

		mockedRepository.DoesExistByCPFFunc = func(ctx context.Context, cpf vos.CPF) (bool, error) {
			return true, nil
		}

		acc := accounts.NewAccount(fakeName, fakeCpf, fakeSecret)
		err := mockedUseCase.CreateAccount(context.Background(), acc)

		assert.Equal(t, accounts.ErrAccountAlreadyExists, err)
	})
}
