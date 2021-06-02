package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app"
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

func (c AccountController) NewAccount(w http.ResponseWriter, r *http.Request) {
	logger := hlog.FromRequest(r)

	var dto dtos.CreateAccountDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		responses.Error(w, logger, app.NewMalformedJSONError())
		return
	}

	if err := c.validator.Struct(dto); err != nil {
		responses.Error(w, logger, err)
		return
	}

	cmd := commands.NewCreateAccountCommand(dto.Secret, dto.CPF, dto.Name)

	_, err := c.bus.Publish(logger.WithContext(r.Context()), cmd)
	if err != nil {
		responses.Error(w, logger, err)
		return
	}

	responses.JSON(w, logger, http.StatusCreated, nil)
}

func (c AccountController) Me(w http.ResponseWriter, r *http.Request) {

}

func (c AccountController) UpdateBalance(w http.ResponseWriter, r *http.Request) {

}
