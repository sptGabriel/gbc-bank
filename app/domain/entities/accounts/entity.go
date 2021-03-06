package accounts

import (
	"github.com/sptGabriel/banking/app/domain/vos"
	"time"
)

const InitialBalance = 0

type Account struct {
	Id        vos.AccountId
	Name      vos.Name
	CPF       vos.CPF
	Secret    vos.Secret
	Balance   int
	CreatedAt time.Time
}

func NewAccount(name vos.Name, cpf vos.CPF, secret vos.Secret) Account {
	return Account{
		Id:      vos.NewAccountId(),
		CPF:     cpf,
		Name:    name,
		Secret:  secret,
		Balance: InitialBalance,
	}
}

func (acc Account) IsEmpty() bool {
	return acc == Account{}
}

func (acc *Account) DebitAmount(amount int) error {
	if acc.Balance-amount < 0 {
		return ErrInsufficientBalance
	}
	acc.Balance -= amount
	return nil
}

func (acc *Account) CreditAmount(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	acc.Balance += amount
	return nil
}
