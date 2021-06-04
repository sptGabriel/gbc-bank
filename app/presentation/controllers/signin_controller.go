package controllers

import (
	"context"
	"encoding/json"
	"github.com/sptGabriel/banking/app/application/dtos"
	"github.com/sptGabriel/banking/app/domain/commands"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"net/http"
)

type signInController struct {
	bus       mediator.Bus
}

type SignInController interface {
	SignIn(r *http.Request) responses.Response
}

func NewSignInController(b mediator.Bus) *signInController {
	return &signInController{bus: b}
}

func (c signInController) SignIn(r *http.Request) responses.Response {
	var dto dtos.SignInDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.IsError(err)
	}

	//if err := c.validator.Struct(dto); err != nil {
	//	return responses.IsError(err)
	//}

	command := commands.NewSignInCommandCommand(dto.CPF, dto.Secret)
	token, err := c.bus.Publish(context.Background(), command)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(token)
}
