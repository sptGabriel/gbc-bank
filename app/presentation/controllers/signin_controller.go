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
// SignIn
// @Description Returns a token to be used on authenticated routes
// @tags login
// @Accept json
// @Produce json
// @Param credentials body dtos.SignInDTO true "Credentials"
// @Success 200 {object} schemas.TokenSchema
// @failure 400 {object} responses.Response
// @failure 409 {object} responses.Response
// @failure 500 {object} responses.Response
// @Router /login [post]
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
