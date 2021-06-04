package middlewares

import (
	"context"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app/application/ports"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"net/http"
	"strings"
)

func AuthHandle(cipherService ports.CipherService,next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := hlog.FromRequest(r)
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			msg := responses.Error{Message: "Malformed token"}
			responses.WriteResponse(w,msg, http.StatusUnauthorized, l)
			return
		}

		jwtToken := authHeader[1]

		claim, err := cipherService.Decrypt(jwtToken)
		if err != nil {
			msg := responses.Error{Message: "Malformed token"}
			responses.WriteResponse(w,msg, http.StatusUnauthorized, l)
			return
		}

		ctx := context.WithValue(r.Context(), "acc_cl", claim)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
