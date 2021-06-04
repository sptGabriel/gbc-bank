package ports

type HashService interface {
	Hash(secret *string) error
	Compare(hashedPassword []byte, password []byte) error
}
