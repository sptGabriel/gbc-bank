package schemas

import (
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
	"time"
)

type AccountSchema struct {
	Id        vos.AccountId    `json:"id"`
	Name      vos.Name    `json:"name"`
	CPF       vos.CPF    `json:"cpf"`
	Balance   int   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccountSchema(account entities.Account) AccountSchema {
	return AccountSchema{
		Id:        account.Id,
		Name:      account.Name,
		CPF:       account.CPF,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}
}