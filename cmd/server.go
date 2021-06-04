package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/handlers"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/common/adapters"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres/account"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres/transfer"
	infraHttp "github.com/sptGabriel/banking/app/infrastructure/http"
	"github.com/sptGabriel/banking/app/infrastructure/http/middlewares"
	"github.com/sptGabriel/banking/app/infrastructure/http/routes"
	"github.com/sptGabriel/banking/app/infrastructure/logger"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

func main() {
	// load application configurations
	cfg := app.ReadConfig(".env")
	// init zero logger instance
	logger := logger.NewLogger()
	// connect to postgres
	conn, err := postgres.ConnectPool(cfg.Postgres.DSN())
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to connect to database")
	}
	defer conn.Close()
	// run postgres migrations
	if err = postgres.RunMigrations(cfg.Postgres.URL()); err != nil {
		logger.Fatal().Err(err).Msg("failed to run postgres migrations")
	}

	// init adapters
	jwtAdapter := adapters.NewJWTAdapter(cfg.Auth.Key, cfg.Auth.Duration)
	bcryptAdapter := adapters.NewBCryptAdapter(10)
	// init repositories
	transactional:= postgres.NewTransactional(conn)
	accountRepository := account.NewRepository(conn)
	transferRepository := transfer.NewRepository(conn)

	// init handlers
	newAccountHandler := handlers.NewCreateAccountHandler(accountRepository, bcryptAdapter)
	makeTransferHandler := handlers.NewMakeTransferHandler(transferRepository, accountRepository, transactional)
	getAccountsHandler := handlers.NewGetAccountsHandler(accountRepository)
	signinHandler:= handlers.NewSignInHandler(accountRepository, jwtAdapter, bcryptAdapter)

	// init command bus
	bus := initBus(logger, newAccountHandler, makeTransferHandler, getAccountsHandler, signinHandler)

	// init controllers
	accountController := controllers.NewAccountController(bus)
	transferController := controllers.NewTransferController(bus)
	signController := controllers.NewSignInController(bus)

	//	init routes
	accRoute := routes.NewAccountRoute(accountController)
	authRoute := routes.NewAuthRouter(signController)
	transferRoute := routes.NewTransferRouter(transferController, jwtAdapter)

	// router
	router := initRouter(transferRoute, accRoute, accRoute, authRoute)

	//	create http server instance
	s := &http.Server{
		Handler:      middlewares.Recovery(router),
		Addr:         fmt.Sprintf("127.0.0.1:%d", cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.ReadTimeout,
		WriteTimeout: cfg.HttpServer.WriteTimeout,
	}
	// run http server
	infraHttp.RunServer(s, logger)
}

func initRouter (routes ...ports.Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	apiRoutes := router.PathPrefix("/api/v1").Subrouter()
	for _, route := range routes {
		route.Init(apiRoutes)
	}
	return  router
}

func initBus (l *zerolog.Logger,  handlers ...ports.Handler) mediator.Bus {
	bus := mediator.NewBus()
	for _, handler := range handlers {
		if err := handler.Init(bus); err != nil {
			l.Fatal().Err(err).Msg("failed to run postgres migrations")
		}
	}
	return bus
}