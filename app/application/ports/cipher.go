package ports

type CipherService interface {
	Encrypt(id string) (string, error)
	Decrypt(val string) (string, error)
}
