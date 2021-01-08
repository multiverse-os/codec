package keyderivation

// "github.com/awnumar/memguard"
// type lockedBuffer = *memguard.LockedBuffer

type Key struct {
	Secret []byte
	Salt   []byte
}

type KeyManager interface {
	Generate(password string, salt []byte) (bool, error)

	Destroy() error

	Key() []byte
	Salt() []byte
}
