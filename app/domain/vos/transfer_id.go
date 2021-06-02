package vos

import "github.com/google/uuid"

type TransferId uuid.UUID

func NewTransferId() TransferId {
	return TransferId(uuid.New())
}

func (id TransferId) String() string {
	return uuid.UUID(id).String()
}
