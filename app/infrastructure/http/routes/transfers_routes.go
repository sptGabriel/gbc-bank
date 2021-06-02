package routes

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/controllers"
	"net/http"
)

type TransferRouter struct {
	controller *controllers.TransferController
}

func NewTransferRouter(b mediator.Bus, v *validator.Validate) *TransferRouter {
	controller := controllers.NewTransferController(b, v)
	return &TransferRouter{controller}
}

func (r *TransferRouter) Init(router *mux.Router) {
	var transfersPath = "transfers"
	router.HandleFunc(fmt.Sprintf("/%s", transfersPath), r.controller.MakeTransfer).Methods(http.MethodPost)
}
