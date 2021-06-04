package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/application/schemas"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type getAccountTransfersHandler struct {
	repository repositories.TransferRepository
}

func NewGetAccountTransferHandler(repository repositories.TransferRepository) *getAccountTransfersHandler {
	return &getAccountTransfersHandler{repository}
}

func (h *getAccountTransfersHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	operation := "Handlers.GetAccountTransfers"

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.GetAccountTransfersCommand)
	if !ok {
		return nil, app.Err(operation, errors.New("invalid transfer command"))
	}

	transfers, err := h.repository.GetAll(ctx, vos.AccountId(command.Id))
	if err != nil {
		return nil, err
	}

	return h.parseToSchema(transfers), err
}

func (h *getAccountTransfersHandler) Init(bus mediator.Bus) error {
	return bus.RegisterHandler(commands.GetAccountTransfersCommand{}, h)
}

func (h *getAccountTransfersHandler) parseToSchema(tr []entities.Transfer) []schemas.TransferSchema {
	transfers := make([]schemas.TransferSchema, 0)

	for _, transfer := range tr {
		schema := schemas.NewTransferSchema(transfer)
		transfers = append(transfers, schema)
	}

	return transfers
}