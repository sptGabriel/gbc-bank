package middlewares

import (
	"context"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"github.com/sptGabriel/banking/app/gateway/ports"
	"net/http"
	"strings"
)

func AuthHandle(cipherService ports.Cipher, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := hlog.FromRequest(r)
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			err := app.ErrMalformedToken
			responses.WriteResponse(w, responses.Error{Message: err.Error()}, http.StatusBadRequest, l)
			return
		}

		jwtToken := authHeader[1]

		claim, err := cipherService.Decrypt(jwtToken)
		if err != nil {
			err := app.ErrUnauthorized
			responses.WriteResponse(w, responses.Error{Message: err.Error()}, http.StatusUnauthorized, l)
			return
		}

		ctx := context.WithValue(r.Context(), "acc_cl", claim)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
