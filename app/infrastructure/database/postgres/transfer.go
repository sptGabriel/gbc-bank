package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type transferRepository struct {
	conn *pgxpool.Pool
}

func NewTransferRepository(c *pgxpool.Pool) repositories.TransferRepository {
	return &transferRepository{c}
}

func (r transferRepository) Create(ctx context.Context, transfer *entities.Transfer) error {
	var query = `
		INSERT INTO
			transfers (id, account_origin_id, account_destination_id, amount)
		VALUES
			($1, $2, $3, $4)
	`
	if _, err := getConnFromCtx(ctx, r.conn).Exec(ctx, query,
		transfer.Id,
		transfer.AccountOriginId,
		transfer.AccountDestinationId,
		transfer.Amount,
	); err != nil {
		return err
	}

	return nil
}

func (r transferRepository) GetAll(ctx context.Context, accountId vos.AccountId) ([]entities.Transfer, error) {
	var query = `
		SELECT
			id, account_origin_id, account_destination_id, amount, created_at
		FROM transfers
		WHERE account_origin_id = $1 OR account_destination_id = $1
		ORDER BY created_at desc
	`
	var transfers []entities.Transfer

	rows, err := r.conn.Query(ctx, query, accountId)
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
