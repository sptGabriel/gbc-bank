package transfer

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/vos"
)

func (r Repository) GetAll(ctx context.Context, accountId vos.AccountId) ([]entities.Transfer, error) {
	var query = `
		SELECT
			id, account_origin_id, account_destination_id, amount, created_at
		FROM transfers
		WHERE account_origin_id = $1 OR account_destination_id = $1
		ORDER BY created_at desc
	`
	var transfers []entities.Transfer

	rows, err := r.pool.Query(ctx, query, accountId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return transfers, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transfer entities.Transfer
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
		transfers = append(transfers, transfer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}
