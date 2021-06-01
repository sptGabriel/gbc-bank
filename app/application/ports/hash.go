package ports

import "github.com/sptGabriel/banking/app/domain/vos"

type Hasher interface {
	Hash(secret *string) error
	Compare(hashed vos.Secret, plainSecret string) error
}
