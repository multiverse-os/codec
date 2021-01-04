package symmetric

////////////////////////////////////////////////////////////////////////////////
type Key struct {
	Secret []byte
}

////////////////////////////////////////////////////////////////////////////////
func (self Key) SecretKey(secretKey []byte) Key {
	self.Secret = secretKey
	return self
}

////////////////////////////////////////////////////////////////////////////////
func (self Key) Cryptosystem() string {
	return "symmetric"
}

func (self Key) Encrypt(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Key) Decrypt(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Key) Sign(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Key) Verify(input []byte) (bool, error) {
	return false, nil
}

////////////////////////////////////////////////////////////////////////////////
