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
	GetAccountTransfers (r *http.Request) responses.Response
}

func NewTransferController(b mediator.Bus) *transferController {
	return &transferController{bus: b}
}

// MakeTransfer
// @Description Do Make a new transfer
// @tags Transfer
// @Accept  json
// @Produce  json
// @Success 201 {object} interface{}
// @failure 400 {object} responses.Error
// @failure 409 {object} responses.Error
// @failure 500 {object} responses.Error
// @Router /api/v1/transfers [post]
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

	return responses.Created(nil)
}

// GetAccountTransfers @Summary Transfer
// @Description Do get all transfers from account
// @Tags Transfer
// @Accept  json
// @Produce  json
// @Success 200 {object} []schemas.TransferSchema
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/transfers [GET]
func (c transferController) GetAccountTransfers(r *http.Request) responses.Response {
	logger := hlog.FromRequest(r)

	accountId := utils.ToUUID(r.Context().Value("acc_cl").(string))
	cmd := commands.NewGetAccountTransfersCommand(accountId)

	res, err := c.bus.Publish(logger.WithContext(r.Context()), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(res)
}
