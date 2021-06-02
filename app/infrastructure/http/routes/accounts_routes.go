package routes

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type AccountsRouter struct {
	controller *controllers.AccountController
}

func NewAccountRouter(b mediator.Bus, v *validator.Validate) *AccountsRouter {
	controller := controllers.NewAccountController(b, v)
	return &AccountsRouter{controller}
}

func (r *AccountsRouter) Init(router *mux.Router) {
	var accountPath = "accounts"
	router.HandleFunc(fmt.Sprintf("/%s", accountPath), r.controller.NewAccount).Methods(http.MethodPost)
}
