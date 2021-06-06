 package schemas

import (
	"github.com/sptGabriel/banking/app/domain/vos"
)

type AccountBalance struct {
	Id        vos.AccountId `json:"id"`
	Balance   int           `json:"balance"`
}

func NewAccountBalanceSchema(Id vos.AccountId, balance int) AccountBalance {
	return AccountBalance{
		Id:       Id,
		Balance:   balance,
	}
}
