package ports

type Hash interface {
	Hash(secret *string) error
	Compare(hashedPassword []byte, password []byte) error
}
