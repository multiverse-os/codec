package symmetric

import (
	"fmt"

	params "github.com/multiverse-os/codec/crypto/params"

	bcrypt "golang.org/x/crypto/bcrypt"
)

////////////////////////////////////////////////////////////////////////////////
//
// **bcrypt** is a key derivation function for passwords. A bcrypt hash string
// is of the form:
//
//      $2b$[cost]$[22 character salt][31 character hash]
//
//  For example:
//
//      $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
//      \__/\/ \____________________/\_____________________________/
//       Alg Cost      Salt                        Hash
//
//  Where:
//    $2a$: The hash algorithm identifier (bcrypt)
//    10: Cost factor (210 ==> 1,024 rounds)
//    N9qo8uLOickgx2ZMRZoMye: 16-byte (128-bit) salt, base64 encoded to 22 characters
//    IjZAgcfl7p92ldGxad68LJZdL17lhWy: 24-byte (192-bit) hash, base64 encoded to 31 characters
//
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// Asymmetric / Symmetric Key Functions
////////////////////////////////////////////////////////////////////////////////
type SymmetricKey interface {
	PasswordHash(password string) ([]byte, error)
	IsPasswordValid(password string) (bool, error)
}

type Key struct {
	secret    []byte
	hash      PasswordHash
	algorithm Algorithm
	params    params.Params
}

////////////////////////////////////////////////////////////////////////////////
func Generate() Key {
	return Key{
		algorithm: Version2b,
		params:    make(params.Params),
	}
}

func (self Key) Password(password []byte) Key {
	self.Secret = password
	return self
}

func (self Key) Hash(hash []byte) Key {

}

////////////////////////////////////////////////////////////////////////////////
func (self Key) Encrypt(plainText []byte) ([]byte, error)  { return []byte{}, nil }
func (self Key) Decrypt(cipherText []byte) ([]byte, error) { return []byte{}, nil }

func (self Key) Sign(input []byte) ([]byte, error) { return []byte{}, nil }
func (self Key) Verify(input []byte) (bool, error) { return false, nil }

////////////////////////////////////////////////////////////////////////////////
// TODO: Store these in params.Parmas
func (self Key) Salt(salt []byte) Key {
	self.Params.AddBytes("salt", salt)
	return self
}

func (self Key) Cost(cost int) Key {
	self.Params.AddInteger("cost", cost)
	return self
}

func (self Key) Algorithm(algorithm Algorithm) Key {
	self.Hash.Algorithm = algorithm
	return self
}

////////////////////////////////////////////////////////////////////////////////
func (self Key) PasswordHash(seed []byte) (Key, error) {
	if hash, err := bcrypt.GenerateFromPassword(seed, bcrypt.DefaultCost); err != nil {
		return Key{}, err
	} else {
		return Key{
			Secret:    seed,
			Algorithm: Version2b,
			Params:    make(params.Params),
			Hash:      hash,
		}, nil
	}
}

func (self Key) IsPasswordValid(password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(self.Hash, []byte(password)); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

////////////////////////////////////////////////////////////////////////////////
type PasswordHash struct {
	Algorithm Algorithm
	Cost      int
	Salt      []byte
	Hash      []byte
}

func (self PasswordHash) Valid() (bool, error) {
	if len(self.Hash) != 31 {
		return false, fmt.Errorf("invalid hash length")
	} else if len(self.Salt) != 22 {
		return false, fmt.Errorf("invalid salt length")
	} else {
		return true, nil
	}
}

func (self PasswordHash) String() string {
	return fmt.Sprintf("$%s$%s$%s%s", self.Algorithm, self.Cost, self.Salt, self.Hash)
}
