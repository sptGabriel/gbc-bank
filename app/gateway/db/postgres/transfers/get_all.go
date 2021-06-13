package transfers

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (r Repository) GetAll(ctx context.Context, accountId vos.AccountId) ([]transfers.Transfer, error) {
	const operation = "Repository.Transfers.GetAll"

	var query = `
		SELECT
			id, account_origin_id, account_destination_id, amount, created_at
		FROM transfers
		WHERE account_origin_id = $1 OR account_destination_id = $1
		ORDER BY created_at desc
	`

	rows, err := r.pool.Query(ctx, query, accountId)
	if err != nil {
		return nil, app.Err(operation, err)
	}
	defer rows.Close()

	var transferList []transfers.Transfer
	for rows.Next() {
		var transfer transfers.Transfer
		err := rows.Scan(
			&transfer.Id,
			&transfer.AccountOriginId,
			&transfer.AccountDestinationId,
			&transfer.Amount,
			&transfer.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transferList = append(transferList, transfer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transferList, nil
}
