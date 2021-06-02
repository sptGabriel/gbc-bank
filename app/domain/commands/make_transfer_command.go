package commands

type MakeTransferCommand struct {
	AccountOriginId      string
	AccountDestinationId string
	Amount               int
}

func NewMakeTransferCommand(accountOriginId string, accountDestinationId string, amount int) MakeTransferCommand {
	return MakeTransferCommand{
		AccountOriginId:      accountOriginId,
		AccountDestinationId: accountDestinationId,
		Amount:               amount,
	}
}
