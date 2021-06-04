package controllers

import (
	"encoding/json"
	"github.com/rs/zerolog/hlog"
	"github.com/sptGabriel/banking/app/application/dtos"
	"github.com/sptGabriel/banking/app/domain/commands"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"github.com/sptGabriel/banking/app/utils"
	"net/http"
)

type transferController struct {
	bus       mediator.Bus
}

type TransferController interface {
	MakeTransfer(r *http.Request) responses.Response
}

func NewTransferController(b mediator.Bus) *transferController {
	return &transferController{bus: b}
}

func (c transferController) MakeTransfer(r *http.Request) responses.Response {
	logger := hlog.FromRequest(r)

	var dto dtos.MakeTransferDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.IsError(err)
	}

	accountOriginID := utils.ToUUID(r.Context().Value("acc_cl").(string))
	accountDestinationID := utils.ToUUID(dto.AccountDestinationId)

	cmd := commands.NewMakeTransferCommand(accountOriginID, accountDestinationID, dto.Amount)

	_, err := c.bus.Publish(logger.WithContext(r.Context()), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(nil)
}
