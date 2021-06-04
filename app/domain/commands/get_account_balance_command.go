package commands

import "github.com/google/uuid"

type GetBalanceCommand struct {
	Id uuid.UUID
}

func NewGetBalanceCommand(Id uuid.UUID) GetBalanceCommand {
	return GetBalanceCommand{Id}
}