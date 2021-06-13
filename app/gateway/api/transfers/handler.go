package transfers

import (
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/domain/usecases/transfers"
	md "github.com/sptGabriel/banking/app/gateway/api/shared/middlewares"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"github.com/sptGabriel/banking/app/gateway/ports"
	"net/http"
)

type Handler interface {
	Create(r *http.Request) responses.Response
	GetTransfers(r *http.Request) responses.Response
}

type handler struct {
	useCase transfers.UseCase
}

func NewHandler(router *mux.Router, useCase transfers.UseCase, cipher ports.Cipher) *handler {
	const path = "/transfers"
	h := &handler{useCase}
	router.HandleFunc(path, md.AuthHandle(cipher, md.RouteAdapter(h.Create))).Methods(http.MethodPost)
	router.HandleFunc(path, md.AuthHandle(cipher, md.RouteAdapter(h.GetTransfers))).Methods(http.MethodGet)
	return h
}
