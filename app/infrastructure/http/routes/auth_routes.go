package routes

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/infrastructure/http/middlewares"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type AuthRouter struct {
	ctrl *controllers.SignInController
}

func NewAuthRouter(b mediator.Bus, v *validator.Validate) *AuthRouter {
	controller := controllers.NewSignInController(b, v)
	return &AuthRouter{controller}
}

func (r *AuthRouter) Init(router *mux.Router) {
	router.HandleFunc("/signin", middlewares.Handle(r.ctrl.SignIn)).Methods(http.MethodPost)
}
