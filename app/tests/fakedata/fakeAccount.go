package fakedata

import (
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func FakeAccount() *entities.Account {
	cpf, _ := vos.NewCpf("10757060099")
	name,_ := vos.NewName("STONE stone")
	secret, _ := vos.NewSecret("stone Stone")
	account := entities.NewAccount(name, cpf, secret)
	account.CreditAmount(2000)
	return &account
}
