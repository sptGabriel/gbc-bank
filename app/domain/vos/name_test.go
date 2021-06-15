package vos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("Should create name vo", func(t *testing.T) {
		validString := "testing name vo"
		name, err := NewName(validString)
		assert.Nil(t, err)
		assert.Equal(t, validString, name.String())
	})
	t.Run("Should fail to create name vo", func(t *testing.T) {
		validString := ""
		name, err := NewName(validString)
		assert.Equal(t, ErrInvalidAccountName, err)
		assert.Equal(t, Name{}, name)
	})
}
