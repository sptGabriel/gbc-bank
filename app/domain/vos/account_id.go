package vos

import (
	"encoding/json"
	"github.com/google/uuid"
)

type AccountId uuid.UUID

func NewAccountId() AccountId {
	return AccountId(uuid.New())
}

func (id AccountId) String() string {
	return uuid.UUID(id).String()
}

func (id AccountId) Equals(accId AccountId) bool {
	return id.String() == accId.String()
}

func (id *AccountId) MarshalJSON() ([]byte, error) {
	byteString, err := json.Marshal(id.String())
	return byteString, err
}
