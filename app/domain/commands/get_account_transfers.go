package commands

import "github.com/google/uuid"

type GetAccountTransfersCommand struct {
	Id uuid.UUID
}

func NewGetAccountTransfersCommand(Id uuid.UUID) GetAccountTransfersCommand {
	return GetAccountTransfersCommand{Id}
}