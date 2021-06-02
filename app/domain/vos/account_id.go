package vos

import "github.com/google/uuid"

type AccountId uuid.UUID

func NewAccountId() AccountId {
	return AccountId(uuid.New())
}

func (id AccountId) String() string{
	return uuid.UUID(id).String()
}