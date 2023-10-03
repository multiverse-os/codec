package symmetric

import (
	"fmt"

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

// //////////////////////////////////////////////////////////////////////////////
// Asymmetric / Symmetric Key Functions
// //////////////////////////////////////////////////////////////////////////////
type SymmetricKey interface {
	PasswordHash(password string) ([]byte, error)
	IsPasswordValid(password string) (bool, error)
}

type Key struct {
	Secret    []byte
	Hash      PasswordHash
	Algorithm Algorithm
	Options   *Options
}

type Options struct {
	Salt []byte
	Cost int
}

// //////////////////////////////////////////////////////////////////////////////
func Generate() Key {
	return Key{
		Algorithm: Version2b,
		Options:   &Options{},
	}
}

func (key Key) Password(password []byte) Key {
	key.Secret = password
	return key
}

func (key Key) Hash(hash []byte) Key {

}

// //////////////////////////////////////////////////////////////////////////////
func (key Key) Encrypt(plainText []byte) ([]byte, error)  { return []byte{}, nil }
func (key Key) Decrypt(cipherText []byte) ([]byte, error) { return []byte{}, nil }

func (key Key) Sign(input []byte) ([]byte, error) { return []byte{}, nil }
func (key Key) Verify(input []byte) (bool, error) { return false, nil }

// //////////////////////////////////////////////////////////////////////////////
// TODO: Store these in params.Parmas
func (key *Key) Salt(salt []byte) Key {
	key.Options.Salt = salt
	return key
}

func (key *Key) Cost(cost int) Key {
	key.Options.Cost = cost
	return key
}

func (key *Key) Algorithm(algorithm Algorithm) Key {
	key.Hash.Algorithm = algorithm
	return key
}

// //////////////////////////////////////////////////////////////////////////////
func (key Key) PasswordHash(seed []byte) (Key, error) {
	if hash, err := bcrypt.GenerateFromPassword(seed, bcrypt.DefaultCost); err != nil {
		return Key{}, err
	} else {
		return Key{
			Secret:    seed,
			Algorithm: Version2b,
			Options:   Options{},
			Hash:      hash,
		}, nil
	}
}

func (key Key) IsPasswordValid(password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(key.Hash, []byte(password)); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// //////////////////////////////////////////////////////////////////////////////
type PasswordHash struct {
	Algorithm Algorithm
	Cost      int
	Salt      []byte
	Hash      []byte
}

func (key PasswordHash) Valid() (bool, error) {
	if len(key.Hash) != 31 {
		return false, fmt.Errorf("invalid hash length")
	} else if len(key.Salt) != 22 {
		return false, fmt.Errorf("invalid salt length")
	} else {
		return true, nil
	}
}

func (key PasswordHash) String() string {
	return fmt.Sprintf("$%s$%s$%s%s", key.Algorithm, key.Cost, key.Salt, key.Hash)
}
