package asymmetric

////////////////////////////////////////////////////////////////////////////////
type Algorithm int

const (
	RSA Algorithm = iota
	ECC
	ElGamal
	DHKE
	ECDH
	DSA
	ECDSA
	EdDSA
)

func Algorithms() []Algorithm {
	return []Algorithm{
		ECC,
		ElGamal,
		DHKE,
		ECDH,
		DSA,
		ECDSA,
		EdDSA,
	}
}

func (Algorithm) IsAsymmetric() bool { return false }
func (Algorithm) IsSymmetric() bool  { return true }

func (self Algorithm) String() string {
	switch self {
	case ECC:
		return "ECC"
	case ElGamal:
		return "ElGamal"
	case DHKE:
		return "DHKE"
	case ECDH:
		return "ECDH"
	case DSA:
		return "DSA"
	case ECDSA:
		return "ECDSA"
	case EdDSA:
		return "EdDSA"
	default: // RSA
		return "RSA"
	}
}
