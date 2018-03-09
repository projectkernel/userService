package auth

type Encryption interface {
	Encrypt(string) string
	Decrypt(string) string
}
