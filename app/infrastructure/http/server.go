package http

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"github.com/sptGabriel/banking/app/application/handlers"
	"github.com/sptGabriel/banking/app/common/adapters"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer(s *http.Server, log *zerolog.Logger) {
	serverErrors := make(chan error, 1)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			serverErrors <- err
		}
	}()
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		sig := <-signals
		log.Info().Msgf("captured signal: %v - server shutdown", sig)
		signal.Stop(signals)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			s.Close()
		}
	}()
	err := <-serverErrors
	if !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err)
	}
}

func InitBus(conn *pgxpool.Pool, bus *mediator.Bus) error {

	// init repositories
	transactionalRepo := postgres.NewTransactional(conn)
	accountRepo := postgres.NewAccountRepository(conn)
	transferRepo := postgres.NewTransferRepository(conn)
	// init handlers
	hasher := adapters.NewBCryptAdapter(10)
	newAccountHandler := handlers.NewCreateAccountHandler(accountRepo, hasher)
	makeTransferHandler := handlers.NewMakeTransferHandler(transferRepo,accountRepo, transactionalRepo)
	// register handlers on the bus
	if err := bus.RegisterHandler(commands.CreateAccountCommand{}, newAccountHandler); err != nil {
		return err
	}
	if err := bus.RegisterHandler(commands.MakeTransferCommand{}, makeTransferHandler); err != nil {
		return err
	}
	return nil
}
