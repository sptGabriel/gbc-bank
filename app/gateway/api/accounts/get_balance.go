package accounts

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
)

type AccountBalance struct {
	Id      vos.AccountId `json:"id"`
	Balance int           `json:"balance"`
}

func NewAccountBalance(Id vos.AccountId, balance int) AccountBalance {
	return AccountBalance{
		Id:      Id,
		Balance: balance,
	}
}

func (h handler) GetBalance(r *http.Request) responses.Response {
	const operation = "Handlers.Accounts.CreateAccount"

	accountId, err := uuid.Parse(mux.Vars(r)["account_id"])
	if err != nil {
		return responses.BadRequest(app.Err(operation, err))
	}

	account, err := h.useCase.GetBalance(r.Context(), vos.AccountId(accountId))
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(NewAccountBalance(account.Id, account.Balance))
}
