package vos

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/sptGabriel/banking/app/utils"
	"strings"
)

type Name struct {
	value string
}

const (
	minNameLength = 8
	maxNameLength = 30
)

var ErrInvalidAccountName = errors.New("account name not valid")

func NewName(name string) (Name, error) {
	if name == "" || !utils.InRange(minNameLength, maxNameLength, len(strings.TrimSpace(name))) {
		return Name{}, ErrInvalidAccountName
	}
	return Name{value: name}, nil
}

func (n Name) String() string {
	return n.value
}

func (n Name) Value() (driver.Value, error) {
	return n.String(), nil
}

func (n *Name) Scan(v interface{}) error {
	if v == nil {
		*n = Name(Name{})
		return nil
	}

	if value, ok := v.(string); ok {
		*n = Name(Name{value})
		return nil
	}

	return errors.New("unable to assign row value to Name")
}

func (n Name) MarshalJSON() ([]byte, error) {
	byteString, err := json.Marshal(n.String())
	return byteString, err
}
