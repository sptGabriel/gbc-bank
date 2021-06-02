package utils

import "github.com/google/uuid"

func ToUUID(id string) uuid.UUID {
	uuid, err := uuid.Parse(id)

	if err != nil {

	}

	return uuid
}
