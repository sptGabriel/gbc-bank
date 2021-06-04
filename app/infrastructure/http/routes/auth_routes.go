package routes

import (
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/infrastructure/http/middlewares"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type authRouter struct {
	ctrl controllers.SignInController
}

func NewAuthRouter(ctrl controllers.SignInController) *authRouter {
	return &authRouter{ctrl}
}

func (r *authRouter) Init(router *mux.Router) {
	router.HandleFunc("/login", middlewares.Handle(r.ctrl.SignIn)).Methods(http.MethodPost)
}
