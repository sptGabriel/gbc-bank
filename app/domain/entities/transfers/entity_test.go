package transfers

import (
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntity(t *testing.T) {

	t.Run("Should create transfer", func(t *testing.T) {
		AccountOriginId := vos.NewAccountId()
		AccountDestinationId := vos.NewAccountId()
		amount := 10000

		res, err := NewTransfer(AccountOriginId, AccountDestinationId, amount)

		assert.Nil(t, err)
		assert.Equal(t, AccountOriginId, res.AccountOriginId)
		assert.Equal(t, AccountDestinationId, res.AccountDestinationId)
		assert.Equal(t, amount, res.Amount)
	})
	t.Run("Should return ErrSELFTransfer", func(t *testing.T) {
		id := vos.NewAccountId()
		amount := 10000

		res, err := NewTransfer(id, id, amount)

		assert.Equal(t, Transfer{}, res)
		assert.Equal(t, ErrSELFTransfer, err)
	})
}
