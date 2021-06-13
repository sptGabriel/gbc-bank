package transfers

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type Repository interface {
	Create(ctx context.Context, transfer *Transfer) error
	GetAll(ctx context.Context, accountId vos.AccountId) ([]Transfer, error)
}
