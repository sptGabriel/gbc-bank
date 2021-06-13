package accounts

import (
	"encoding/json"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
)

type CreateAccountRequest struct {
	Name   string `json:"name" validate:"required,min=10"`
	CPF    string `json:"cpf" validate:"required,min=9,max=12"`
	Secret string `json:"secret" validate:"required,min=8"`
}

func (h handler) Create(r *http.Request) responses.Response {
	const operation = "Handlers.Accounts.CreateAccount"

	var dto CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.BadRequest(app.Err(operation, err))
	}
	// create vos
	cpf, err := vos.NewCPF(dto.CPF)
	if err != nil {
		return responses.BadRequest(err)
	}
	name, err := vos.NewName(dto.Name)
	if err != nil {
		return responses.BadRequest(err)
	}
	secret, err := vos.NewSecret(dto.Secret)
	if err != nil {
		return responses.BadRequest(err)
	}

	account := accounts.NewAccount(name, cpf, secret)

	if err := h.useCase.CreateAccount(r.Context(), account); err != nil {
		return responses.IsError(err)
	}

	return responses.Created(nil)
}
