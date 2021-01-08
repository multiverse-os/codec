package crypto

import (
	asymmetric "github.com/multiverse-os/codec/crypto/key/asymmetric"
	params "github.com/multiverse-os/codec/crypto/key/params"
	symmetric "github.com/multiverse-os/codec/crypto/key/symmetric"
)

////////////////////////////////////////////////////////////////////////////////
const (
	RSA   = asymmetric.RSA
	ECC   = asymmetric.ECC
	ECDH  = asymmetric.ECDH
	DSA   = asymmetric.DSA
	ECDSA = asymmetric.ECDSA
	EdDSA = asymmetric.EdDSA
)

const (
	Argon2    = symmetric.Argon2
	BCrypt    = symmetric.BCrypt
	PBKDF2    = symmetric.PBKDF2
	SCrypt    = symmetric.SCrypt
	DES       = symmetric.DES
	TripleDES = symmetric.TripleDES
	AES       = symmetric.AES
	Blowfish  = symmetric.Blowfish
	Twofish   = symmetric.Twofish
	ChaCha20  = symmetric.ChaCha20
)

////////////////////////////////////////////////////////////////////////////////
type Algorithm interface {
	String() string
	IsSymmetric() bool
	IsAsymmetric() bool
	DefaultParams(p params.Params) params.Params
}

func Algorithms() (algorithms []Algorithm) {
	for _, a := range symmetric.Algorithms() {
		algorithms = append(algorithms, a)
	}
	for _, a := range asymmetric.Algorithms() {
		algorithms = append(algorithms, a)
	}
	return algorithms
}
