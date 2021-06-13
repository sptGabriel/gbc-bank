package accounts

import "errors"

var (
	ErrBalanceUpdate        = errors.New("could not update balance from account")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrInvalidAccountID     = errors.New("account id not valid")
	ErrInsufficientBalance  = errors.New("account has insufficient balance")
	ErrInvalidAmount        = errors.New("could not credit this amount")
)
