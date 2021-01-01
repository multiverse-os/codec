package otp

//GenerateKey(seed string) Keypair
//GenerateRandomKey() Keypair
type Type int

const (
	TOTP Type = iota
	HOTP
)

// Aliasing
const (
	Time = TOTP
	HMAC = HOTP
)

type OneTimePassword struct {
	Type         Type
	SharedSecret string
	Host         string
	Username     string
}

func (self OneTimePassword) URI() string {
	return ""
}
