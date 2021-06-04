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
	Create(r *http.Request) responses.Response
	GetAccounts(r *http.Request) responses.Response
	GetBalance(r *http.Request) responses.Response
}

func NewAccountController(b mediator.Bus) *accountController {
	return &accountController{bus: b}
}

// Create @Summary accounts
// @Description Do create a new account
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param Body body dtos.CreateAccountDTO true "Body"
// @Success 201 {object} interface{}
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts  [POST]
func (c accountController) Create(r *http.Request) responses.Response {
	var dto dtos.CreateAccountDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.IsError(err)
	}

	cmd := commands.NewCreateAccountCommand(dto.Secret, dto.CPF, dto.Name)

	_, err := c.bus.Publish(context.Background(), cmd)
	if err != nil {
		return responses.IsError(err)
	}

	return responses.Created(nil)
}


// GetAccounts @Summary accounts
// @Description Do get all accounts
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} []schemas.AccountSchema
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts [GET]
func (c accountController) GetAccounts(r *http.Request) responses.Response {
	accounts, err := c.bus.Publish(context.Background(), commands.GetAllAccountsCommand{})
	if err != nil {
		return responses.IsError(err)
	}

	return responses.OK(accounts)
}

// GetBalance @Summary accounts
// @Description Do get account balance
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} []schemas.AccountBalance
// @Failure 404 {object} responses.Error
// @Failure 422 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts/{account_id}/balance [GET]
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
