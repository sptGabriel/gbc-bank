package fakedata

import (
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func FakeTransfer() *entities.Transfer {
	destinationId := vos.NewAccountId()
	currentId := vos.NewAccountId()
	transfer, _ := entities.NewTransfer(currentId, destinationId, 1)
	return &transfer
}
