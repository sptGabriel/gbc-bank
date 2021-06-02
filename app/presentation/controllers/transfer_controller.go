package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app/application/dtos"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"github.com/sptGabriel/banking/app/utils"
	"net/http"
)

type TransferController struct {
	bus       mediator.Bus
	validator *validator.Validate
}

func NewTransferController(b mediator.Bus, v *validator.Validate) *TransferController {
	return &TransferController{bus: b, validator: v}
}

func (c TransferController) MakeTransfer(r *http.Request) responses.Response {
	logger := hlog.FromRequest(r)

	var dto dtos.MakeTransferDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.IsError(err)
	}

	if err := c.validator.Struct(dto); err != nil {
		return responses.IsError(err)
	}

	accountOriginID := utils.ToUUID(dto.AccountOriginId)
	accountDestinationID := utils.ToUUID(dto.AccountDestinationId)

	cmd := commands.NewMakeTransferCommand(accountOriginID, accountDestinationID, dto.Amount)

	_, err := c.bus.Publish(logger.WithContext(r.Context()), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(nil)
}
