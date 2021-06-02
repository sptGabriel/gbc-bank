package middlewares

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"net/http"
)

func Handle(handler func(r *http.Request) responses.Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := hlog.FromRequest(r)

		res := handler(r)
		if res.Error != nil {
			logger.Error().Err(res.Error)
		}

		WriteResponse(w, res.Data, res.Status, logger)
	}
}

func WriteResponse(w http.ResponseWriter, d interface{}, st int, l *zerolog.Logger) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(st)
	if d != nil {
		if err := json.NewEncoder(w).Encode(d); err != nil {
			l.Error().Stack().Err(err).Msg("failed to encode response")
		}
	}
}