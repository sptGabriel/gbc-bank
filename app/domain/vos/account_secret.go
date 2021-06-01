package vos

import (
	"github.com/sptGabriel/banking/app"
)

const secretMinLength = 8

var ErrAccountSecretInvalid = app.NewDomainError("invalid secret")

type Secret struct {
	value string
}

func NewSecret(secret string) (Secret, error) {
	if secret == "" || len(secret) < secretMinLength {
		return Secret{}, ErrAccountSecretInvalid
	}
	return Secret{value: secret}, nil
}

func (secret Secret) String() string {
	return secret.value
}
