package asymmetric

import (
	params "github.com/multiverse-os/codec/crypto/key/params"
)

type PublicKey []byte
type PrivateKey []byte

type Permissions struct {
	Sign   bool
	Subkey bool
}

type Contact struct {
	Name    string
	Email   string
	Comment string
}

type Keypair struct {
	Algorithm Algorithm

	PublicKey  PublicKey
	PrivateKey PrivateKey

	Params params.Params

	Contact
	Permissions

	// NOTE: Symmetrically encrypt the private key
	Password string

	Parent  *Keypair
	Subkeys []*Keypair
}

func (self Keypair) KeyAlgorithm() Algorithm { return self.Algorithm }

func (self Keypair) String() string { return self.Algorithm.String() }

func (self Keypair) Encrypt(plainText []byte) ([]byte, error)  { return []byte{}, nil }
func (self Keypair) Decrypt(cipherText []byte) ([]byte, error) { return []byte{}, nil }

func (self Keypair) Sign(input []byte) ([]byte, error) { return []byte{}, nil }
func (self Keypair) Verify(input []byte) (bool, error) { return false, nil }
