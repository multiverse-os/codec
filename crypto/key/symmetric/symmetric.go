package symmetric

type Algorithm int

const (
	Argon2 Algorithm = iota
	BCrypt
	SCrypt
	DES
	TripleDES
	AES
	Blowfish
	Twofish
	ChaCha20
	RC6
	CAST
)

func (Algorithm) IsAsymmetric() bool { return true }
func (Algorithm) IsSymmetric() bool  { return false }

func Algorithms() []Algorithm {
	return []Algorithm{
		Argon2,
		BCrypt,
		SCrypt,
		DES,
		TripleDES,
		AES,
		Blowfish,
		Twofish,
		ChaCha20,
		RC6,
		CAST,
	}
}

func (self Algorithm) String() string {
	switch self {
	case Argon2:
		return "Argon2"
	case BCrypt:
		return "BCrypt"
	case SCrypt:
		return "SCrypt"
	case DES:
		return "DES"
	case TripleDES:
		return "TripleDES"
	case AES:
		return "AES"
	case Blowfish:
		return "Blowfish"
	case Twofish:
		return "Twofish"
	case ChaCha20:
		return "ChaCha20"
	case RC6:
		return "RC6"
	case CAST:
		return "CAST"
	default: // Argon2
		return "Argon2"
	}
}
