package rsa

////////////////////////////////////////////////////////////////////////////////
type Keypair struct {
	privateKey []byte
	publicKey  []byte
}

////////////////////////////////////////////////////////////////////////////////
func (self Keypair) PublicKey(publicKey []byte) Keypair {
	self.publicKey = publicKey
	return self
}

func (self Keypair) PrivateKey(privateKey []byte) Keypair {
	self.privateKey = privateKey
	return self
}

////////////////////////////////////////////////////////////////////////////////
func (self Keypair) Cryptosystem() string {
	return "asymmetric"
}

func (self Keypair) Encrypt(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Keypair) Decrypt(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Keypair) Sign(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Keypair) Verify(input []byte) (bool, error) {
	return false, nil
}
