package commands

type CreateAccountCommand struct {
	Secret string
	Cpf    string
	Name   string
}

func NewCreateAccountCommand(secret string, cpf string, name string) CreateAccountCommand {
	return CreateAccountCommand{
		Secret: secret,
		Cpf:    cpf,
		Name:   name,
	}
}
