package accounts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/services"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/api/shared/middlewares"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
)

type Handler interface {
	Authenticate(r *http.Request) responses.Response
}

type handler struct {
	authService services.Authenticate
}

type AuthenticateRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
type AuthenticateResponse struct {
	Token string `json:"token"`
}

func NewAuthenticateResponse(token string) AuthenticateResponse {
	return AuthenticateResponse{token}
}

func NewHandler(router *mux.Router, authService services.Authenticate) *handler {
	h := &handler{authService}
	router.HandleFunc("/login", middlewares.RouteAdapter(h.Authenticate)).Methods(http.MethodPost)
	return h
}

// Authenticate
// @Description Returns a token to be used on authenticated routes
// @tags login
// @Accept json
// @Produce json
// @Param credentials body AuthenticateRequest true "Credentials"
// @Success 200 {object} AuthenticateResponse
// @failure 400 {object} responses.Response
// @failure 409 {object} responses.Response
// @failure 500 {object} responses.Response
// @Router /login [post]
func (h handler) Authenticate(r *http.Request) responses.Response {
	const operation = "Handlers.Authentication.Authenticate"

	var dto AuthenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.BadRequest(app.Err(operation, err))
	}

	cpf, err := vos.NewCPF(dto.CPF)
	if err != nil {
		return responses.BadRequest(err)
	}

	secret, err := vos.NewSecret(dto.Secret)
	if err != nil {
		return responses.BadRequest(err)
	}

	res, err := h.authService.Authenticate(r.Context(), cpf, secret)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(NewAuthenticateResponse(res))
}
