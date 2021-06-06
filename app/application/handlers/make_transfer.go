package handlers

import (
	"context"
	"errors"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"

	"github.com/sptGabriel/banking/app/infrastructure/database/postgres"
	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"time"
)

type makeTransferHandler struct {
	transferRepo  repositories.TransferRepository
	accountRepo   repositories.AccountRepository
	transactional postgres.Transactional
}

func NewMakeTransferHandler(t repositories.TransferRepository,
	ar repositories.AccountRepository,
	tr postgres.Transactional,
) *makeTransferHandler {
	return &makeTransferHandler{
		transferRepo:  t,
		accountRepo:   ar,
		transactional: tr,
	}
}

func (h *makeTransferHandler) Execute(ctx context.Context, cmd mediator.Command) (interface{}, error) {
	operation := "Handlers.MakeTransfer"

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	command, ok := cmd.(commands.MakeTransferCommand)
	if !ok {
		return nil, app.Err(operation, errors.New("invalid make transfer command"))
	}

	originAccountId := vos.AccountId(command.AccountOriginId)

	destinationId := vos.AccountId(command.AccountDestinationId)

	transfer, err := entities.NewTransfer(originAccountId, destinationId, command.Amount)
	if err != nil {
		return nil, err
	}

	res, err := h.transactional.Exec(ctx, func(txCtx context.Context) (interface{}, error) {
		if err := h.debitOriginAccount(txCtx, originAccountId, transfer.Amount); err != nil {
			return nil, err
		}
		if err := h.creditTargetAccount(txCtx, destinationId, transfer.Amount); err != nil {
			return nil, err
		}
		err := h.transferRepo.Create(txCtx, &transfer)
		return nil, err
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *makeTransferHandler) debitOriginAccount(ctx context.Context, id vos.AccountId, amount int) error {
	account, err := h.accountRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := account.DebitAmount(amount); err != nil {
		return err
	}

	return h.accountRepo.UpdateBalance(ctx, account)
}

func (h *makeTransferHandler) creditTargetAccount(ctx context.Context, id vos.AccountId, amount int) error {
	account, err := h.accountRepo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if err := account.CreditAmount(amount); err != nil {
		return err
	}

	return h.accountRepo.UpdateBalance(ctx, account)
}

func (h *makeTransferHandler) Init(bus mediator.Bus) error {
	return bus.RegisterHandler(commands.MakeTransferCommand{}, h)
}