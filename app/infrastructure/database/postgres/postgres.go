package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/sptGabriel/banking/app"
)

func ConnectPool(url string) (*pgxpool.Pool, error) {
	operation:= "Postgres.ConnectPool"

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, app.Err(operation,err)
	}

	config.ConnConfig.Logger = zerologadapter.NewLogger(log.Logger)
	dbPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, app.Err(operation,err)
	}

	return dbPool, nil
}
