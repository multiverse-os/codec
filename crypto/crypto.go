package crypto

import "github.com/multiverse-os/starshipyard/framework/datastore/leveldb/codec/crypto/key"

type Type int

const (
	Assymetric Type = iota
	Symmetric
)

func (self Type) String() string {
	switch self {
	case Symmetric:
		return "symetric"
	default: // Assymetric
		return "assymetric"
	}
}

func Cipher(cipherType Type, algorithm Algorithm) key.Cipher {
	return key.Cipher{Type: cipherType, Algorithm: algorithm}
}

//type CryptographicKey interface {
//	GenerateChildKey() key.Pair
//	// NOTE: Derive a SSH, Bitcoin or other key deterministically (incrementable)
//	//       from any key type, for example: Use a PGP key to generate your SSH
//	//       keys on all your remote machines.
//	DeriveKeypair(algorithm Algorithm) Keypair
//
//	OneTimePassword(otpType OTPType) []byte
//
//	ExpiresAt(expiresAt time.Time) Keypair
//	Name(name string) Keypair
//	ContactURI(contact string) Keypair
//	Password(symmetricEncryption string) Keypair
//
//	Algorithm() string
//
//	Encrypt(input []byte) ([]byte, error)
//	Decrypt(input []byte) ([]byte, error)
//
//	Sign(input []byte) ([]byte, error)
//	VerifySignature(input []byte) (bool, error)
//
//	SplitKey() []KeyPart
//	AssembleKey(keyParts ...KeyPart) Keypair
//
//	GenerateToken(expiresAt time.Time) Token
//	GenerateCertificate(expiresAt time.Time) Certificate
//}
//
//func GenerateKey(algorithm AlgorithmType, seed []byte) (Keypair, error) {
//	switch algorithm {
//	case Argon2:
//		keypair := Keypair{
//			Algorithm: Algorithm{
//				Type:   algorithm,
//				Params: argon2.DefaultParams(),
//			},
//		}
//
//		return keypair
//
//	default:
//		return Keypair{}, fmt.Errorf("invalid algorithm")
//	}
//}
