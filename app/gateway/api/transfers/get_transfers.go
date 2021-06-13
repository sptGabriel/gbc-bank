package transfers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities/transfers"
	"github.com/sptGabriel/banking/app/domain/vos"
	"github.com/sptGabriel/banking/app/gateway/api/shared/responses"
	"net/http"
	"time"
)

type GetTransfersResponse struct {
	TransferId    vos.TransferId `json:"id"`
	DestinationId vos.AccountId  `json:"account_destination_id"`
	Amount        int            `json:"amount"`
	CreatedAt     time.Time      `json:"created_at"`
}

func NewGetTransferResponse(transfer transfers.Transfer) GetTransfersResponse {
	return GetTransfersResponse{
		TransferId:    transfer.Id,
		DestinationId: transfer.AccountDestinationId,
		Amount:        transfer.Amount,
		CreatedAt:     transfer.CreatedAt,
	}
}

func (h handler) GetTransfers(r *http.Request) responses.Response {
	const operation = "Handlers.Transfers.GetTransfers"

	accountId, err := uuid.Parse(r.Context().Value("acc_cl").(string))
	if err != nil {
		return responses.BadRequest(app.Err(operation, fmt.Errorf("invalid params: account id")))
	}

	transfers, err := h.useCase.GetTransfers(r.Context(), vos.AccountId(accountId))
	if err != nil {
		return responses.IsError(err)
	}

	output := make([]GetTransfersResponse, 0)

	for _, transfer := range transfers {
		output = append(output, NewGetTransferResponse(transfer))
	}

	return responses.OK(output)
}
