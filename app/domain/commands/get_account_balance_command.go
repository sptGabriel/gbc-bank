package commands

import "github.com/sptGabriel/banking/app/domain/vos"

type GetAllAccountBalanceCommand struct {
	Id vos.AccountId
}
