package app

import (
	"errors"
	"fmt"
)

var (
	ErrInternal       = errors.New("internal error")
	ErrUnauthorized   = errors.New("unauthorized error")
	ErrMalformedToken = errors.New("token is not valid")
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
