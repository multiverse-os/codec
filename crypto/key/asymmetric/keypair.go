package asymmetric

type Permissions struct {
	Sign   bool
	Subkey bool
}

type PrivateKey struct {
	Key        []byte
	Fingerpint []byte
}

type PublicKey struct {
	Key        []byte
	Fingerpint []byte
}

type Contact struct {
	Name    string
	Email   string
	Comment string
}

type Keypair struct {
	PublicKey  PublicKey
	PrivateKey PrivateKey

	Contact
	Permissions

	// NOTE: Symmetrically encrypt the private key
	Password string

	Parent  *Keypair
	Subkeys []*Keypair
}

func (self Keypair) Encrypt(plainText []byte) ([]byte, error)  { return []byte{}, nil }
func (self Keypair) Decrypt(cipherText []byte) ([]byte, error) { return []byte{}, nil }

func (self Keypair) Sign(input []byte) ([]byte, error) { return []byte{}, nil }
func (self Keypair) Verify(input []byte) (bool, error) { return false, nil }
