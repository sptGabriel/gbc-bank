package entities

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

func (account Account) IsEmpty() bool {
	return account == Account{}
}
