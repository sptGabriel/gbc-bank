package repositories

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type TransferRepository interface {
	Create(ctx context.Context, account *entities.Transfer) error
	GetAll(ctx context.Context, accountId vos.AccountId) ([]entities.Transfer, error)
}
