package dtos

type MakeTransferDTO struct {
	//AccountOriginId      string `json:"account_origin_id" validate:"required,uuid"`
	AccountDestinationId string `json:"account_destination_id" validate:"required,uuid"`
	Amount               int    `json:"amount" validate:"required"`
}
