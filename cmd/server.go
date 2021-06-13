package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/common/adapters"
	"github.com/sptGabriel/banking/app/domain/services"
	accUseCases "github.com/sptGabriel/banking/app/domain/usecases/accounts"
	transferUseCases "github.com/sptGabriel/banking/app/domain/usecases/transfers"
	accHandler "github.com/sptGabriel/banking/app/gateway/api/accounts"
	authHandler "github.com/sptGabriel/banking/app/gateway/api/authentication"
	"github.com/sptGabriel/banking/app/gateway/api/shared/middlewares"
	transferHandler "github.com/sptGabriel/banking/app/gateway/api/transfers"
	"github.com/sptGabriel/banking/app/gateway/db/postgres"
	"github.com/sptGabriel/banking/app/gateway/db/postgres/accounts"
	"github.com/sptGabriel/banking/app/gateway/db/postgres/transfers"
	"github.com/sptGabriel/banking/app/gateway/logger"
	swagger "github.com/sptGabriel/banking/docs/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// repositories
	transferRepo := transfers.NewRepository(conn)
	accRepo := accounts.NewRepository(conn)
	transactional := postgres.NewTransactional(conn)

	// init domain services
	authService := services.NewAuth(accRepo, bcryptAdapter, jwtAdapter)

	// useCases
	accountsUseCase := accUseCases.NewUseCase(accRepo, bcryptAdapter)
	transfersUseCase := transferUseCases.NewUseCase(accRepo, transferRepo, transactional)

	// configure global swagger with environment variable
	swagger.SwaggerInfo.Host = cfg.Swagger.SwaggerHost

	// init router
	router := mux.NewRouter()
	router.PathPrefix("/docs/v1/swagger").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
	apiRoutes := router.PathPrefix("/api/v1").Subrouter()
	router.NotFoundHandler = middlewares.NotFoundHandle()
	router.Use(middlewares.CORSHandle)
	router.Use(middlewares.JsonHandle)

	// init router handlers
	accHandler.NewHandler(apiRoutes, accountsUseCase)
	transferHandler.NewHandler(apiRoutes, transfersUseCase, jwtAdapter)
	authHandler.NewHandler(apiRoutes, authService)

	// create http server instance
	s := &http.Server{
		Handler:      middlewares.Recovery(router),
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.ReadTimeout,
		WriteTimeout: cfg.HttpServer.WriteTimeout,
	}

	// run http server
	RunServer(s, logger)
}

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
		fmt.Println(err)
		log.Error().Err(err)
	}
}
