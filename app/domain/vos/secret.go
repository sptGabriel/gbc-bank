package vos

import (
	"database/sql/driver"
	"errors"
	"github.com/sptGabriel/banking/app/gateway/ports"
)

const secretMinLength = 8

type Secret struct {
	value    string
	isHashed bool
}

var ErrInvalidAccountSecret = errors.New("account secret not valid")

func NewSecret(secret string) (Secret, error) {
	if secret == "" || len(secret) < secretMinLength {
		return Secret{}, ErrInvalidAccountSecret
	}
	return Secret{value: secret, isHashed: false}, nil
}

func (s *Secret) Encrypt(hash ports.Hash) error {
	if s.isHashed {
		return nil
	}
	err := hash.Hash(&s.value)
	if err == nil {
		s.isHashed = true
	}
	return err
}

func (s Secret) String() string {
	return s.value
}

func (s Secret) IsHashed() bool {
	return s.isHashed
}

func (s Secret) Value() (driver.Value, error) {
	return s.String(), nil
}

func (s *Secret) Scan(v interface{}) error {
	if v == nil {
		*s = Secret(Secret{})
		return nil
	}

	if value, ok := v.(string); ok {
		*s = Secret(Secret{value, true})
		return nil
	}

	return errors.New("unable to assign row value to Secret")
}
