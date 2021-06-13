package transfers

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/gateway/db/postgres"
)

func (r Repository) Create(ctx context.Context, transfer *transfers.Transfer) error {
	var query = `
		INSERT INTO
			transfers (id, account_origin_id, account_destination_id, amount)
		VALUES
			($1, $2, $3, $4)
	`
	if _, err := postgres.GetConnFromCtx(ctx, r.pool).Exec(ctx, query,
		transfer.Id,
		transfer.AccountOriginId,
		transfer.AccountDestinationId,
		transfer.Amount,
	); err != nil {
		return err
	}

	return nil
}
