package vos

import (
	"encoding/json"
	"github.com/google/uuid"
)

type TransferId uuid.UUID

func NewTransferId() TransferId {
	return TransferId(uuid.New())
}

func (id TransferId) String() string {
	return uuid.UUID(id).String()
}
func (id TransferId) MarshalJSON() ([]byte, error) {
	byteString, err := json.Marshal(id.String())
	return byteString, err
}
