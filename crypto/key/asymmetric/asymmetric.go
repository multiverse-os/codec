package asymmetric

import (
	params "github.com/multiverse-os/codec/crypto/key/params"
)

////////////////////////////////////////////////////////////////////////////////
type KeyType int

const (
	Public KeyType = iota
	Private
)

////////////////////////////////////////////////////////////////////////////////
type Algorithm int

const (
	RSA Algorithm = iota
	ECC
	ECDH
	DSA
	ECDSA
	EdDSA
)

func Algorithms() []Algorithm {
	return []Algorithm{
		ECC,
		ECDH,
		DSA,
		ECDSA,
		EdDSA,
	}
}

func (Algorithm) IsAsymmetric() bool { return false }
func (Algorithm) IsSymmetric() bool  { return true }

func (self Algorithm) String() string {
	switch self {
	case ECC:
		return "ECC"
	case ECDH:
		return "ECDH"
	case DSA:
		return "DSA"
	case ECDSA:
		return "ECDSA"
	case EdDSA:
		return "EdDSA"
	default: // RSA
		return "RSA"
	}
}

func (Algorithm) DefaultParams(p params.Params) params.Params {
	return p
}
