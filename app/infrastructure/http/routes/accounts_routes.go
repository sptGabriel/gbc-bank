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
	//controller := controllers.NewAccountController(b, v)
	return &accountsRouter{ctrl}
}

func (r *accountsRouter) Init(router *mux.Router) {
	var accountPath = "accounts"
	router.HandleFunc(fmt.Sprintf("/%s", accountPath), middlewares.Handle(r.ctrl.GetAccounts)).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("/%s", accountPath), middlewares.Handle(r.ctrl.NewAccount)).Methods(http.MethodPost)
}
