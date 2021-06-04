package vos

import (
	"database/sql/driver"
	"errors"
	"github.com/sptGabriel/banking/app"
)

const secretMinLength = 8

type Secret struct {
	value string
}

func NewSecret(secret string) (Secret, error) {
	if secret == "" || len(secret) < secretMinLength {
		return Secret{}, app.ErrInvalidAccountSecret
	}
	return Secret{value: secret}, nil
}

func (s Secret) String() string {
	return s.value
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
		*s = Secret(Secret{value})
		return nil
	}

	return errors.New("unable to assign row value to Secret")
}
