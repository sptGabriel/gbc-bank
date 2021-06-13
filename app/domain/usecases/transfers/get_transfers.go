package transfers

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (uc useCase) GetTransfers(ctx context.Context, id vos.AccountId) ([]transfers.Transfer, error) {
	transfers, err := uc.transferRepo.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}
