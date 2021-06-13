package transfers

import (
	"encoding/json"
	"fmt"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"github.com/sptGabriel/banking/app/utils"
	"net/http"
)

type createTransferRequest struct {
	AccountDestinationId string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}

func (h handler) Create(r *http.Request) responses.Response {
	const operation = "Handlers.Transfers.Create"

	var dto createTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return responses.BadRequest(app.Err(operation, err))
	}

	accountOriginID := vos.AccountId(utils.ToUUID(r.Context().Value("acc_cl").(string)))
	accountDestinationID := vos.AccountId(utils.ToUUID(dto.AccountDestinationId))

	transfer, err := transfers.NewTransfer(accountOriginID, accountDestinationID, dto.Amount)
	if err != nil {
		return responses.BadRequest(app.Err(operation, err))
	}

	if err := h.useCase.CreateTransfer(r.Context(), transfer); err != nil {
		fmt.Println(err)
		return responses.IsError(err)
	}

	return responses.Created(nil)
}
