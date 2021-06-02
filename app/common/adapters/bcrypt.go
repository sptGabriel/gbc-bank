package adapters

import (
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/vos"
	"golang.org/x/crypto/bcrypt"
)

type BCryptAdapter struct {
	salt int
}

func NewBCryptAdapter(salt int) BCryptAdapter {
	return BCryptAdapter{
		salt: salt,
	}
}

func (bc BCryptAdapter) Hash(plainText *string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(*plainText), bcrypt.DefaultCost)
	if err != nil {
		return app.NewInternalError("failed to hash password", err)
	}
	*plainText = string(hashed)
	return nil
}

func (bc BCryptAdapter) Compare(secret vos.Secret, plainSecret string) error {
	return bcrypt.CompareHashAndPassword([]byte(secret.String()), []byte(plainSecret))
}
