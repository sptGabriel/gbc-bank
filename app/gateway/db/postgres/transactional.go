package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/rs/zerolog/log"
)

type transactionKey struct{}

type transactional struct {
	conn *pgxpool.Pool
}

type Transactional interface {
	Exec(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error)
}

func NewTransactional(c *pgxpool.Pool) transactional {
	return transactional{
		conn: c,
	}
}

func (t transactional) Exec(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := t.conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("postgres.Transactional: failed to start transaction %w", err)
	}

	defer func() {
		if err == nil {
			err = tx.Commit(ctx)
		}
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			err = fmt.Errorf("postgres.WithTransaction: failed to rollback (%s) %w", rollbackErr.Error(), err)
		}
	}()

	_, err = tx.Exec(ctx, "SET TRANSACTION ISOLATION LEVEL READ COMMITTED")
	if err != nil {
		return nil, err
	}

	ctxTx := context.WithValue(ctx, transactionKey{}, tx)
	data, err := f(ctxTx)
	return data, err
}

func GetConnFromCtx(ctx context.Context, db *pgxpool.Pool) pgxtype.Querier {
	tx, ok := ctx.Value(transactionKey{}).(pgxtype.Querier)
	if !ok {
		return pgxtype.Querier(db)
	}

	return tx
}
