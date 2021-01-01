package key

type Cipher struct {
	Type      Type
	Algorithm Algorithm
}

func (self Cipher) Encrypt(key, input []byte) ([]byte, error) {
	return self.Key(self.Type, key).Encrypt(input)
}

func (self Cipher) Decrypt(key, input []byte) ([]byte, error) {
	return self.Key(self.Type, key).Decrypt(input)
}

func (self Cipher) Sign(privateKey, input []byte) ([]byte, error) {
	return self.PrivateKey(privateKey).Sign(input)
}

func (self Cipher) VerifySignature(publicKey, input []byte) (bool, error) {
	return self.PublicKey(publicKey).VerifySignature(input)
}

func (self Cipher) PrivateKey(key []byte) Keypair {
	return key.Assymetric(key)
}

func (self Cipher) PublicKey(key []byte) Keypair {
	return key.AssymetricKey().PublicKey(key)
}

func (self Cipher) Key(t Type, key []byte) Key {
	switch t {
	case Symmetric:
		return key.Symmetric(key)
	default: // Assymetric
		return key.Assymetric(key)
	}
}
