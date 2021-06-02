package commands

type SignInCommand struct {
	Cpf    string
	Secret string
}

func NewSignInCommandCommand(cpf string, secret string) SignInCommand {
	return SignInCommand{
		Secret: secret,
		Cpf:    cpf,
	}
}
