package vos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCPF(t *testing.T) {
	t.Run("Should create cpf vo", func(t *testing.T) {
		validString := "29891275000"
		cpf, err := NewCPF(validString)
		assert.Nil(t, err)
		assert.Equal(t, validString, cpf.String())
	})
	t.Run("Should fail to create cpf vo", func(t *testing.T) {
		validString := ""
		cpf, err := NewCPF(validString)
		assert.Equal(t, ErrInvalidAccountCPF, err)
		assert.Equal(t, CPF{}, cpf)
	})
}
