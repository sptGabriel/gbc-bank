package vos

import (
	"github.com/sptGabriel/banking/app"
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

var ErrAccountNameInvalid = app.NewDomainError("invalid name")

func NewName(name string) (Name, error) {
	if name == "" || !utils.InRange(minNameLength, maxNameLength, len(strings.TrimSpace(name))) {
		return Name{}, ErrAccountNameInvalid
	}
	return Name{value: name}, nil
}

func (name Name) String() string {
	return name.value
}
