package handlers

import (
	"context"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/utils"
	"time"
)

type MakeTransferHandler struct {
	transferRepo  repositories.TransferRepository
	accountRepo   repositories.AccountRepository
	transactional postgres.Transactional
}

func NewMakeTransferHandler(t repositories.TransferRepository, ar repositories.AccountRepository, tr postgres.Transactional) *MakeTransferHandler {
	return &MakeTransferHandler{
		transferRepo:  t,
		accountRepo:   ar,
		transactional: tr,
	}
}

func (h *MakeTransferHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.MakeTransferCommand)
	if !ok {
		return nil, app.NewInternalError("invalid command", nil)
	}

	originAccountId := vos.AccountId(utils.ToUUID(command.AccountOriginId))

	destinationId := vos.AccountId(utils.ToUUID(command.AccountDestinationId))

	transfer := entities.NewTransfer(originAccountId, destinationId, command.Amount)

	res, err := h.transactional.Exec(ctx, func(txCtx context.Context) (interface{}, error) {
		if err := h.debitOriginAccount(txCtx, originAccountId, transfer.Amount); err != nil {
			return nil, err
		}
		if err := h.debitDestinationAccount(txCtx, destinationId, transfer.Amount); err != nil {
			return nil, err
		}
		err := h.transferRepo.Create(txCtx, &transfer)
		return nil, err
	})

	return res, err
}

func (h *MakeTransferHandler) debitOriginAccount(ctx context.Context, id vos.AccountId, amount int) error {
	account, err := h.accountRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := account.DebitAmount(amount); err != nil {
		return err
	}

	return h.accountRepo.UpdateAccountBalance(ctx, account)
}

func (h *MakeTransferHandler) debitDestinationAccount(ctx context.Context, id vos.AccountId, amount int) error {
	account, err := h.accountRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := account.CreditAmount(amount); err != nil {
		return err
	}

	return h.accountRepo.UpdateAccountBalance(ctx, account)
}
