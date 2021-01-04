package symmetric

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

////////////////////////////////////////////////////////////////////////////////
// Asymmetric / Symmetric Key Functions
////////////////////////////////////////////////////////////////////////////////
type Key struct {
	Salt          []byte
	LastUpdatedAt time.Time
	MinimumLength int
	MaximumLength int
}

func (self Key) Encrypt(plainText []byte) ([]byte, error)  { return []byte{}, nil }
func (self Key) Decrypt(cipherText []byte) ([]byte, error) { return []byte{}, nil }

func (self Key) Sign(input []byte) ([]byte, error) { return []byte{}, nil }
func (self Key) Verify(input []byte) (bool, error) { return false, nil }

////////////////////////////////////////////////////////////////////////////////
// Symmetric Key Functions
////////////////////////////////////////////////////////////////////////////////
type SymmetricKey struct {
	//Hash(password string) (
}

type CryptoInt interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

func (crp *Crypto) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (crp *Crypto) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
