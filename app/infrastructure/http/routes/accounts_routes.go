package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/infrastructure/http/middlewares"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type accountsRouter struct {
	ctrl controllers.AccountController
}

func NewAccountRoute(ctrl controllers.AccountController) *accountsRouter {
	return &accountsRouter{ctrl}
}

func (r *accountsRouter) Init(router *mux.Router) {
	var accountPath = "/accounts"
	var balancePath = fmt.Sprintf("%s/{account_id}/balance", accountPath)
	router.HandleFunc(accountPath, middlewares.Handle(r.ctrl.GetAccounts)).Methods(http.MethodGet)
	router.HandleFunc(accountPath, middlewares.Handle(r.ctrl.Create)).Methods(http.MethodPost)
	router.HandleFunc(balancePath, middlewares.Handle(r.ctrl.GetBalance)).Methods(http.MethodGet)
}
