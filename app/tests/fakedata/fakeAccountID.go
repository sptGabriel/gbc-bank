package fakedata

import (
	"github.com/google/uuid"
	"github.com/sptGabriel/banking/app/utils"
)

func FakeAccountID () uuid.UUID {
	id := "abce8b02-5f3a-4f2c-96a7-964e37d0dc08"
	return utils.ToUUID(id)
}
