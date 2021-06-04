package adapters

import (
	"github.com/sptGabriel/banking/app"
	"golang.org/x/crypto/bcrypt"
)

type bCryptAdapter struct {
	salt int
}

func NewBCryptAdapter(salt int) bCryptAdapter {
	return bCryptAdapter{
		salt: salt,
	}
}

func (bc bCryptAdapter) Hash(plainText *string) error {
	operation := "Adapters.BCrypt.Hash"

	hashed, err := bcrypt.GenerateFromPassword([]byte(*plainText), bcrypt.DefaultCost)
	if err != nil {
		return app.Err(operation, err)
	}
	*plainText = string(hashed)
	return nil
}

func (bc bCryptAdapter) Compare(hashedPassword []byte, password []byte) error {
	operation := "Adapters.BCrypt.Compare"

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err == nil {
		return nil
	}

	return app.Err(operation, err)
}
