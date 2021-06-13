package ports

type Cipher interface {
	Encrypt(id string) (string, error)
	Decrypt(val string) (string, error)
}
