package crypto

type Algorithm int

const (
	Argon2 Algorithm = iota
	BCrypt
	SCrypt
	None
)

func (self Algorithm) String() string {
	switch self {
	case Argon2:
		return "argon2"
	case BCrypt:
		return "bcrypt"
	case SCrypt:
		return "scrypt"
	default: // None
		return "none"
	}
}
