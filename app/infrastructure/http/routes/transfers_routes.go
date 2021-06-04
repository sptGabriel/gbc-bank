package routes

import (
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/application/ports"
	md "github.com/sptGabriel/banking/app/infrastructure/http/middlewares"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type transferRouter struct {
	ctrl       controllers.TransferController
	cipherService ports.CipherService
}

func NewTransferRouter(ctrl controllers.TransferController, cipherService ports.CipherService) *transferRouter {
	return &transferRouter{ctrl, cipherService}
}

func (r *transferRouter) Init(router *mux.Router) {
	var path = "/transfers"
	router.HandleFunc(path, md.AuthHandle(r.cipherService, md.Handle(r.ctrl.MakeTransfer))).Methods(http.MethodPost)
	router.HandleFunc(path, md.AuthHandle(r.cipherService, md.Handle(r.ctrl.GetAccountTransfers))).Methods(http.MethodGet)
}
