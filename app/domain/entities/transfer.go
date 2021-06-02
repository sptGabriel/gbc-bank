package entities

import (
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/vos"
	"time"
)

type Transfer struct {
	Id                   vos.TransferId
	AccountOriginId      vos.AccountId
	AccountDestinationId vos.AccountId
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(accountOriginId vos.AccountId, accountDestinationId vos.AccountId, amount int) (Transfer, error) {
	if accountOriginId.Equals(accountDestinationId) {
		return Transfer{}, app.NewDomainError("the origin account cannot be the same as the destination account")
	}
	return Transfer{
		Id:                   vos.NewTransferId(),
		AccountOriginId:      accountOriginId,
		AccountDestinationId: accountDestinationId,
		Amount:               amount,
	}, nil
}
