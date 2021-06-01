package vos

import (
	"github.com/Nhanderu/brdoc"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/utils"
)

type CPF struct {
	value string
}

var ErrAccountCpfInvalid = app.NewDomainError("invalid cpf")

func NewCpf(cpf string) (CPF, error) {
	if cpf := brdoc.IsCPF(cpf); !cpf {
		return CPF{}, ErrAccountCpfInvalid
	}
	return CPF{utils.Clean(cpf)}, nil
}

func (c CPF) String() string {
	return c.value
}
