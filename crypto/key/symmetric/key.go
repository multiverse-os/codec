package symmetric

import (
	"crypto/rand"
	"fmt"
	"io"
)

////////////////////////////////////////////////////////////////////////////////
func (self Algorithm) SaltLength() int {
	switch self {
	case Argon2:
		return 32
	case BCrypt:
		return 32
	case PBKDF2:
		return 32
	default:
		return 32
	}
}

func (self Algorithm) HashLength() int {
	switch self {
	case Argon2:
		return 64
	case BCrypt:
		return 64
	case PBKDF2:
		return 64
	default:
		return 64
	}
}

////////////////////////////////////////////////////////////////////////////////
type Variant int

////////////////////////////////////////////////////////////////////////////////
type PasswordHash struct {
	Variant Variant
	Cost    int
	Salt    []byte
	Hash    []byte
}

func (self PasswordHash) Valid() (bool, error) { return true, nil }
func (self PasswordHash) String() string       { return fmt.Sprintf("") }

func (self PasswordHash) RandomSalt() []byte {
	salt := make([]byte, len(self.Salt))
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic("error reading from random source: " + err.Error())
	}
	return salt
}

////////////////////////////////////////////////////////////////////////////////
// Symmetric Key Functions
////////////////////////////////////////////////////////////////////////////////
type SymmetricKey interface {
	PasswordHash(password string) ([]byte, error)
	IsPasswordValid(password string) (bool, error)
}
