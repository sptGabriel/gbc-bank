package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sptGabriel/banking/app/application/dtos"
	"github.com/sptGabriel/banking/app/domain/commands"
	"github.com/sptGabriel/banking/app/utils"

	"github.com/sptGabriel/banking/app/infrastructure/mediator"
	"github.com/sptGabriel/banking/app/presentation/responses"
	"net/http"
)

type accountController struct {
	bus       mediator.Bus
}

type AccountController interface {
	NewAccount(r *http.Request) responses.Response
	GetAccounts(r *http.Request) responses.Response
	GetBalance(r *http.Request) responses.Response
}

func NewAccountController(b mediator.Bus) *accountController {
	return &accountController{bus: b}
}

func (c accountController) NewAccount(r *http.Request) responses.Response {
	var dto dtos.CreateAccountDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.IsError(err)
	}

	cmd := commands.NewCreateAccountCommand(dto.Secret, dto.CPF, dto.Name)

	_, err := c.bus.Publish(context.Background(), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(nil)
}

func (c accountController) GetAccounts(r *http.Request) responses.Response {
	accounts, err := c.bus.Publish(context.Background(), commands.GetAllAccountsCommand{})
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(accounts)
}

func (c accountController) GetBalance(r *http.Request) responses.Response {
	params := mux.Vars(r)

	accountId := utils.ToUUID(params["account_id"])

	cmd := commands.NewGetBalanceCommand(accountId)

	res, err := c.bus.Publish(context.Background(), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(res)
}
