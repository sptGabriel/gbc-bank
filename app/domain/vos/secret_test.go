package vos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecret(t *testing.T) {
	t.Run("Should create secret vo", func(t *testing.T) {
		validString := "123456789"
		secret, err := NewSecret(validString)
		assert.Nil(t, err)
		assert.Equal(t, validString, secret.String())
	})
	t.Run("Should fail to create secret vo", func(t *testing.T) {
		validString := ""
		secret, err := NewSecret(validString)
		assert.Equal(t, ErrInvalidAccountSecret, err)
		assert.Equal(t, Secret{}, secret)
	})
}