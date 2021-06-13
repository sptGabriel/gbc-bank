package vos

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/Nhanderu/brdoc"
	"github.com/sptGabriel/banking/app/utils"
)

type CPF struct {
	value string
}

var ErrInvalidAccountCPF    = errors.New("account cpf not valid")

func NewCPF(cpf string) (CPF, error) {
	if cpf := brdoc.IsCPF(cpf); !cpf {
		return CPF{}, ErrInvalidAccountCPF
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

func (c CPF) MarshalJSON() ([]byte, error) {
	byteString, err := json.Marshal(c.String())
	return byteString, err
}
