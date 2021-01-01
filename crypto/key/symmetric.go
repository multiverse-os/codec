package key

type Key struct {
	Cipher
	Secret []byte

	Salt []byte

	Hash []byte
}
