package vos

import (
	"database/sql/driver"
	"errors"
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

func (c CPF) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c *CPF) Scan(v interface{}) error {
	if v == nil {
		*c = CPF(CPF{})
		return nil
	}

	if value, ok := v.(string); ok {
		*c = CPF(CPF{value})
		return nil
	}

	return errors.New("unable to assign row value to CPF")
}