package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
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
		fmt.Println(err)
		log.Error().Err(err)
	}
}