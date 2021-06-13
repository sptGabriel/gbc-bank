package accounts

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/domain/usecases/accounts"
	"github.com/sptGabriel/banking/app/gateway/api/shared/middlewares"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
)

type Handler interface {
	Create(r *http.Request) responses.Response
	GetAll(r *http.Request) responses.Response
	GetBalance(r *http.Request) responses.Response
}

type handler struct {
	useCase accounts.UseCase
}

func NewHandler(router *mux.Router, useCase accounts.UseCase) *handler {
	var accountPath = "/accounts"
	var balancePath = fmt.Sprintf("%s/{account_id}/balance", accountPath)
	h := &handler{useCase}
	router.HandleFunc(accountPath, middlewares.RouteAdapter(h.GetAll)).Methods(http.MethodGet)
	router.HandleFunc(accountPath, middlewares.RouteAdapter(h.Create)).Methods(http.MethodPost)
	router.HandleFunc(balancePath, middlewares.RouteAdapter(h.GetBalance)).Methods(http.MethodGet)
	return h
}
