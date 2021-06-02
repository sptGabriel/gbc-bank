package responses

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/sptGabriel/banking/app"
	"net/http"
)

type jsonError struct {
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, l *zerolog.Logger, e error) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(app.GetErrorCode(e))
	if err := json.NewEncoder(w).Encode(jsonError{e.Error()}); err != nil {
		l.Error().Stack().Err(err).Msg("failed to encode error")
	}
}

func JSON(w http.ResponseWriter, l *zerolog.Logger, c int, d interface{}) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(c)
	if d != nil {
		if err := json.NewEncoder(w).Encode(d); err != nil {
			l.Error().Stack().Err(err).Msg("failed to encode response")
		}
	}
}
