package symmetric

type Algorithm int

const (
	Argon2 Algorithm = iota
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
		DES,
		TripleDES,
		Blowfish,
		Twofish,
		ChaCha20,
		RC6,
		CAST,
	}
}

func (self Algorithm) String() string {
	switch self {
	case DES:
		return "DES"
	case TripleDES:
		return "TripleDES"
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
