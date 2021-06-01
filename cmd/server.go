package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
	infraHttp "github.com/sptGabriel/banking/app/infrastructure/http"
	"github.com/sptGabriel/banking/app/infrastructure/http/routes"
	"github.com/sptGabriel/banking/app/infrastructure/logger"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"net/http"
)

func main() {
	cfg := app.ReadConfig(".env")

	logger := logger.NewLogger()

	conn, err := postgres.ConnectPool(cfg.Postgres.DSN())
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to connect to database")
	}
	defer conn.Close()

	if err = postgres.RunMigrations(cfg.Postgres.URL()); err != nil {
		logger.Fatal().Err(err).Msg("failed to run postgres migrations")
	}

	bus := mediator.NewBus()
	if err := infraHttp.InitBus(conn, &bus); err != nil {
		logger.Fatal().Err(err).Msg("error to init command bus")
	}

	validator := validator.New()
	router := mux.NewRouter()
	accRoute := routes.NewAccountRouter(bus, validator, logger)
	accRoute.Init(router)

	s := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.ReadTimeout,
		WriteTimeout: cfg.HttpServer.WriteTimeout,
	}

	infraHttp.RunServer(s, logger)
}
