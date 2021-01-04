package symmetric

type Key struct{}

func (self Key) Encrypt(plainText []byte) ([]byte, error)  { return []byte{}, nil }
func (self Key) Decrypt(cipherText []byte) ([]byte, error) { return []byte{}, nil }

func (self Key) Sign(input []byte) ([]byte, error) { return []byte{}, nil }
func (self Key) Verify(input []byte) (bool, error) { return false, nil }
