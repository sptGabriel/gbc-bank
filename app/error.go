package app

import (
	"errors"
	"fmt"
)

var (
	ErrInternal                = errors.New("internal error")
	ErrAccountAlreadyExists    = errors.New("account already exists")
	ErrBalanceUpdate           = errors.New("could not update balance from account")
	ErrAccountNotFound         = errors.New("account not found")
	ErrInvalidAccountID        = errors.New("account id not valid")
	ErrInvalidAccountName      = errors.New("account name not valid")
	ErrInvalidAccountSecret    = errors.New("account secret not valid")
	ErrInvalidAccountCPF       = errors.New("account cpf not valid")
	ErrInsufficientBalance     = errors.New("account has insufficient balance")
	ErrInvalidAmount           = errors.New("could not credit this amount")
	ErrTransferNotFound        = errors.New("transfer not found")
	ErrAccountTransferNotFound = errors.New("account has no transfers")
	ErrSELFTransfer            = errors.New("origin account cannot be the same as the destination account")
)

type DomainError struct {
	op     string
	err    error
	parent *DomainError
}

func Err(op string, err error) *DomainError {
	if err == nil {
		return nil
	}

	domainError := DomainError{op: op, err: err}

	if error, ok := err.(*DomainError); ok {
		error.parent = &domainError
	}

	return &domainError
}

func (e DomainError) Error() string {
	return fmt.Sprintf(e.err.Error())
}
