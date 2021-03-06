package responses

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/services"
	transfersUseCase "github.com/sptGabriel/banking/app/domain/usecases/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
	accountsRepo "github.com/sptGabriel/banking/app/gateway/db/postgres/accounts"
	"net/http"
)

type Response struct {
	Status int
	Error  error
	Data   interface{}
}

type Error struct {
	Message string `json:"message"`
}

func Created(d interface{}) Response {
	return Response{
		Status: http.StatusCreated,
		Data:   d,
	}
}

func OK(d interface{}) Response {
	return Response{
		Status: http.StatusOK,
		Data:   d,
	}
}

func Updated(d interface{}) Response {
	return Response{
		Status: http.StatusNoContent,
		Data:   d,
	}
}

func Conflict(err error) Response {
	return genericError(http.StatusConflict, err)
}

func NotFound(err error) Response {
	return genericError(http.StatusNotFound, err)
}

func BadRequest(err error) Response {
	return genericError(http.StatusBadRequest, err)
}

func Unauthorized(err error) Response {
	return genericError(http.StatusUnauthorized, err)
}

func genericError(status int, err error) Response {
	return Response{
		Status: status,
		Error:  err,
		Data:   Error{Message: err.Error()},
	}
}

func InternalError(err error) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  err,
		Data:   Error{Message: app.ErrInternal.Error()},
	}
}

func IsError(err error) Response {
	switch {
	case errors.Is(err, accounts.ErrAccountAlreadyExists):
		return Conflict(err)
	case errors.Is(err, accountsRepo.ErrAccountNotFound):
		return NotFound(err)
	case errors.Is(err, transfersUseCase.ErrAccountDestinationNotFound):
		return NotFound(err)
	case errors.Is(err, transfersUseCase.ErrAccountOriginNotFound):
		return NotFound(err)
	case errors.Is(err, accounts.ErrInsufficientBalance):
		return BadRequest(err)
	case errors.Is(err, accounts.ErrBalanceUpdate):
		return BadRequest(err)
	case errors.Is(err, vos.ErrInvalidAccountCPF):
		return BadRequest(err)
	case errors.Is(err, accounts.ErrInvalidAccountID):
		return BadRequest(err)
	case errors.Is(err, vos.ErrInvalidAccountName):
		return BadRequest(err)
	case errors.Is(err, vos.ErrInvalidAccountSecret):
		return BadRequest(err)
	case errors.Is(err, transfers.ErrSELFTransfer):
		return BadRequest(err)
	case errors.Is(err, accounts.ErrInvalidAmount):
		return BadRequest(err)
	case errors.Is(err, services.ErrInvalidCredentials):
		return BadRequest(err)
	default:
		return InternalError(err)
	}
}

func WriteResponse(w http.ResponseWriter, d interface{}, st int, l *zerolog.Logger) {
	w.WriteHeader(st)
	if d != nil {
		if err := json.NewEncoder(w).Encode(d); err != nil {
			l.Error().Stack().Err(err).Msg("failed to encode response")
		}
	}
}
