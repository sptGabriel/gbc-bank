package schemas

import (
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
	"time"
)

type TransferSchema struct {
	TransferId    vos.TransferId `json:"id"`
	DestinationId vos.AccountId `json:"account_destination_id"`
	Amount        int           `json:"amount"`
	CreatedAt     time.Time     `json:"created_at"`
}

func NewTransferSchema(transfer entities.Transfer) TransferSchema {
	return TransferSchema{
		TransferId: transfer.Id,
		DestinationId: transfer.AccountDestinationId,
		Amount:        transfer.Amount,
		CreatedAt:     transfer.CreatedAt,
	}
}
