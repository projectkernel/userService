package auth

type JWT struct {
	secret string
	duration int
}

func NewJWT(secret string, duration int) *JWT {
	return &JWT{
		secret: secret,
		duration: duration,
	}
}

func (jwt JWT) Encrypt(string) string {
	return ""
}
func (jwt JWT) Decrypt(string) string {
	return ""
}