package middlewares

import (
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
)

func RouteAdapter(handler func(r *http.Request) responses.Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := hlog.FromRequest(r)

		res := handler(r)
		if res.Error != nil {
			logger.Error().Err(res.Error)
		}

		responses.WriteResponse(w, res.Data, res.Status, logger)
	}
}
