package key

import (
	"crypto/subtle"

	argon2 "golang.org/x/crypto/argon2"
)

type Keypair struct {
	Cipher

	Salt []byte

	PrivateKey []byte
	PublicKey  []byte

	Hash []byte // Sometimes called Address, also provides merkle hash data

	RootKey   *Keypair
	ParentKey *Keypair
	ChildKeys []*Keypair
}

func (self Keypair) Params(name string) int {
	return self.Algorithm.Parameters[name]
}

func (self Keypair) GeneratePublicKey() []byte {
	switch self.Algorithm.Type {
	case Argon2:
		return argon2.IDKey(
			self.PrivateKey,
			self.Salt,
			uint32(self.Params("iterations")),
			uint32(self.Params("memory")),
			uint8(self.Params("threads")),
			uint32(self.Params("length")),
		)
	default:
		return []byte{}
	}
}

func (self Keypair) IsPrivateKey(seed []byte) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey(seed, salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}