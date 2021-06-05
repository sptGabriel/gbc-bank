package main

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/go-swagger/go-swagger"
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
	swagger "github.com/sptGabriel/banking/docs"
	"net/http"
)

// @host localhost:8080
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
	transactional := postgres.NewTransactional(conn)
	accountRepository := account.NewRepository(conn)
	transferRepository := transfer.NewRepository(conn)

	// init handlers
	newAccountHandler := handlers.NewCreateAccountHandler(accountRepository, bcryptAdapter)
	getAccountsHandler := handlers.NewGetAccountsHandler(accountRepository)
	getBalanceHandler := handlers.NewGetAccountBalanceHandler(accountRepository)
	getAccountTransfers := handlers.NewGetAccountTransferHandler(transferRepository)
	makeTransferHandler := handlers.NewMakeTransferHandler(transferRepository, accountRepository, transactional)
	signinHandler := handlers.NewSignInHandler(accountRepository, jwtAdapter, bcryptAdapter)

	// init command bus
	bus := initBus(
		logger,
		newAccountHandler,
		makeTransferHandler,
		getAccountsHandler,
		signinHandler,
		getBalanceHandler,
		getAccountTransfers,
	)

	// init controllers
	accountController := controllers.NewAccountController(bus)
	transferController := controllers.NewTransferController(bus)
	signController := controllers.NewSignInController(bus)

	//	init routes
	accRoute := routes.NewAccountRoute(accountController)
	authRoute := routes.NewAuthRouter(signController)
	transferRoute := routes.NewTransferRouter(transferController, jwtAdapter)

	// configure global swagger with environment variable
	swagger.SwaggerInfo.Host = cfg.Swagger.SwaggerHost

	// router
	router := initRouter(
		transferRoute,
		accRoute,
		accRoute,
		authRoute,
	)

	//	create http server instance
	s := &http.Server{
		Handler:      middlewares.Recovery(router),
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.ReadTimeout,
		WriteTimeout: cfg.HttpServer.WriteTimeout,
	}
	// run http server
	infraHttp.RunServer(s, logger)
}

func initRouter(routes ...ports.Route) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/docs/v1/swagger").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
	apiRoutes := router.PathPrefix("/api/v1").Subrouter()
	for _, route := range routes {
		route.Init(apiRoutes)
	}
	return router
}

func initBus(l *zerolog.Logger, handlers ...ports.Handler) mediator.Bus {
	bus := mediator.NewBus()
	for _, handler := range handlers {
		if err := handler.Init(bus); err != nil {
			l.Fatal().Err(err).Msg("failed to run postgres migrations")
		}
	}
	return bus
}
