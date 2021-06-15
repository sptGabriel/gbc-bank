package accounts

import (
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntity(t *testing.T) {
	fakeCpf, _ := vos.NewCPF("70405270062")
	fakeName, _ := vos.NewName("testing create account")
	fakeSecret, _ := vos.NewSecret("123456789")
	fakeAccount := NewAccount(fakeName, fakeCpf, fakeSecret)

	t.Run("Should debit amount from account", func(t *testing.T) {
		account := fakeAccount
		account.Balance = 10
		account.DebitAmount(1)
		assert.Equal(t,9, account.Balance)
	})

	t.Run("Should return insufficient balance", func(t *testing.T) {
		account := fakeAccount
		account.Balance = 10
		err := account.DebitAmount(11)
		assert.Equal(t, ErrInsufficientBalance, err)
	})

	t.Run("Should credit amount to account", func(t *testing.T) {
		account := fakeAccount
		expected := 10000
		err := account.CreditAmount(expected)
		assert.Nil(t, err)
		assert.Equal(t, expected, account.Balance)
	})

	t.Run("Should return err invalid amount", func(t *testing.T) {
		account := fakeAccount
		expected := -1
		err := account.CreditAmount(expected)
		assert.Equal(t, ErrInvalidAmount, err)
	})

	t.Run("account empty", func(t *testing.T) {
		account := Account{}
		expected := true
		res := account.IsEmpty()
		assert.Equal(t, expected, res)
	})
}