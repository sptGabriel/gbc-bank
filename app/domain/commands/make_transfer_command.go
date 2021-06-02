package commands

import "github.com/google/uuid"

type MakeTransferCommand struct {
	AccountOriginId      uuid.UUID
	AccountDestinationId uuid.UUID
	Amount               int
}

func NewMakeTransferCommand(accountOriginId uuid.UUID, accountDestinationId uuid.UUID, amount int) MakeTransferCommand {
	return MakeTransferCommand{
		AccountOriginId:      accountOriginId,
		AccountDestinationId: accountDestinationId,
		Amount:               amount,
	}
}
