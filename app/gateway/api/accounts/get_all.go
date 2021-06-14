package accounts

import (
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
	"time"
)

type GetAccountsResponse struct {
	Id        vos.AccountId `json:"id"`
	Name      vos.Name      `json:"name"`
	CPF       vos.CPF       `json:"cpf"`
	Balance   int           `json:"balance"`
	CreatedAt time.Time     `json:"created_at"`
}

func NewGetAccountsResponse(account accounts.Account) GetAccountsResponse {
	return GetAccountsResponse{
		Id:        account.Id,
		Name:      account.Name,
		CPF:       account.CPF,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}
}

// GetAll @Summary accounts
// @Description Do get all accounts
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} []GetAccountsResponse
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts [GET]
func (h handler) GetAll(r *http.Request) responses.Response {
	const operation = "Handlers.Accounts.GetAll"
	accounts, err := h.useCase.GetAll(r.Context())
	if err != nil {
		return responses.IsError(err)
	}

	output := make([]GetAccountsResponse, 0)

	for _, account := range accounts {
		output = append(output, NewGetAccountsResponse(account))
	}

	return responses.OK(output)
}
