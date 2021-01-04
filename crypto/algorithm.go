package crypto

////////////////////////////////////////////////////////////////////////////////
// Components of cryptosystems
////////////////////////////////////////////////////////////////////////////////
// A basic cryptosystem includes the following components:
//
//   * **Plaintext** the data that needs to be protected.
//   * **Ciphertext** the encrypted, or unreadable, version of the plaintext.
//
//   * **Encryption algorithm** the mathematical algorithm that takes plaintext
//     as the input and returns ciphertext. It also produces the unique
//     encryption key for that text.
//   * **Decryption algorithm** the mathematical algorithm that takes
//     ciphertext as the input and decodes it into plaintext. It also uses
//     the unique decryption key for that text.
//
//   * **Encryption key** the value known to the sender that is used to compute
//     the ciphertext for the given plaintext.
//     **Decryption key** the value known to the receiver that is used to
//     decode the given ciphertext into plaintext.
//
////////////////////////////////////////////////////////////////////////////////
type Algorithm interface {
	String() string
	IsAsymmetric() bool
	IsSymmetric() bool
}

////////////////////////////////////////////////////////////////////////////////
type Cryptosystem int

func Algorithms() (algorithms []Algorithm) {
	for _, aa := range AsymmetricAlgorithms() {
		algorithms = append(algorithms, aa)
	}
	for _, aa := range SymmetricAlgorithms() {
		algorithms = append(algorithms, aa)
	}
	return algorithms
}

func (self Cryptosystem) String() string {
	for _, a := range Algorithms() {
		if len(self.String()) == len(a.String()) &&
			self.String() == a.String() {
			a.String()
		}
	}
	return ""
}

func (self Cryptosystem) IsAsymmetric() bool {
	for _, cs := range AsymmetricAlgorithms() {
		if len(self.String()) == len(cs.String()) &&
			self.String() == cs.String() {
			return true
		}
	}
	return false
}

func (self Cryptosystem) IsSymmetric() bool {
	for _, cs := range SymmetricAlgorithms() {
		if len(self.String()) == len(cs.String()) &&
			self.String() == cs.String() {
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
type Symmetric Cryptosystem

const (
	Argon2 Symmetric = iota
	DES
	TripleDES
	AES
	Blowfish
	Twofish
	ChaCha20
	RC6
	CAST
)

func SymmetricAlgorithms() []Symmetric {
	return []Symmetric{
		DES,
		TripleDES,
		Blowfish,
		Twofish,
		ChaCha20,
		RC6,
		CAST,
	}
}

func (Symmetric) IsAsymmetric() bool { return true }
func (Symmetric) IsSymmetric() bool  { return false }

func (self Symmetric) String() string {
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

////////////////////////////////////////////////////////////////////////////////
type Asymmetric Cryptosystem

const (
	RSA Asymmetric = iota
	ECC
	ElGamal
	DHKE
	ECDH
	DSA
	ECDSA
	EdDSA
)

func AsymmetricAlgorithms() []Asymmetric {
	return []Asymmetric{
		ECC,
		ElGamal,
		DHKE,
		ECDH,
		DSA,
		ECDSA,
		EdDSA,
	}
}

func (Asymmetric) IsAsymmetric() bool { return false }
func (Asymmetric) IsSymmetric() bool  { return true }

func (self Asymmetric) String() string {
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

////////////////////////////////////////////////////////////////////////////////
