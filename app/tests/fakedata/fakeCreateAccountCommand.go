package fakedata

import "github.com/sptGabriel/banking/app/domain/commands"

func FakeCreateAccountCMD() commands.CreateAccountCommand {
	return commands.NewCreateAccountCommand("12345678", "10757060099", "HU4HUHU3HUHUHEU")
}
