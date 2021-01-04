package argon2

import (
	params "github.com/multiverse-os/codec/crypto/params"
)

//GenerateKey(seed string) Keypair
//GenerateRandomKey() Keypair
//type Functionality interface {
//	GenerateChildKey() Keypair
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
//}

type Keypair struct {
	PublicKey []byte
}

func GenerateKey(seed []byte) Keypair {
	return Keypair{}
}

func DefaultParams(params params.Value) params.Value {
	params.AddInteger("memory", 16384)
	params.AddInteger("iterators", 3)
	params.AddInteger("threads", 2)
	params.AddInteger("length", 32)
}
