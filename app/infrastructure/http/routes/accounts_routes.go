package routes

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/infrastructure/http/middlewares"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type AccountsRouter struct {
	ctrl *controllers.AccountController
}

func NewAccountRouter(b mediator.Bus, v *validator.Validate) *AccountsRouter {
	controller := controllers.NewAccountController(b, v)
	return &AccountsRouter{controller}
}

func (r *AccountsRouter) Init(router *mux.Router) {
	var accountPath = "accounts"
	router.HandleFunc(fmt.Sprintf("/%s", accountPath), middlewares.Handle(r.ctrl.GetAccounts)).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("/%s", accountPath), middlewares.Handle(r.ctrl.NewAccount)).Methods(http.MethodPost)
}
