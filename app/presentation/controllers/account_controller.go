package controllers

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app/application/dtos"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"net/http"
)

type AccountController struct {
	bus       mediator.Bus
	validator *validator.Validate
}

func NewAccountController(b mediator.Bus, v *validator.Validate) *AccountController {
	return &AccountController{bus: b, validator: v}
}

func (c AccountController) NewAccount(r *http.Request) responses.Response {
	logger := hlog.FromRequest(r)

	var dto dtos.CreateAccountDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.IsError(err)
	}

	if err := c.validator.Struct(dto); err != nil {
		return responses.IsError(err)
	}

	cmd := commands.NewCreateAccountCommand(dto.Secret, dto.CPF, dto.Name)

	_, err := c.bus.Publish(logger.WithContext(r.Context()), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(nil)
}

func (c AccountController) GetAccounts(r *http.Request) responses.Response {
	accounts, err := c.bus.Publish(context.Background(), commands.GetAllAccountsCommand{})
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(accounts)
}
