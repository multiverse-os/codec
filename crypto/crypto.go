package crypto

import (
	asymmetric "github.com/multiverse-os/codec/crypto/key/asymmetric"
	symmetric "github.com/multiverse-os/codec/crypto/key/symmetric"
)

////////////////////////////////////////////////////////////////////////////////
type Keys struct {
	AsymmetricKey asymmetric.Keypair
	SymmetricKey  symmetric.Key
}

////////////////////////////////////////////////////////////////////////////////
func EmptyKeyring() Keys {
	return Keys{
		AsymmetricKey: asymmetric.Keypair{},
		SymmetricKey:  symmetric.Key{},
	}
}

////////////////////////////////////////////////////////////////////////////////
func (self Keys) AddPublicKey(a Algorithm, key []byte) Keys {
	if a.IsAsymmetric() {
		self.AsymmetricKey = asymmetric.Keypair{
			PublicKey: asymmetric.PublicKey(key),
			Algorithm: a.(asymmetric.Algorithm),
		}
	}
	return self
}

func (self Keys) AddPrivateKey(a Algorithm, key []byte) Keys {
	if a.IsAsymmetric() {
		self.AsymmetricKey = asymmetric.Keypair{
			PrivateKey: asymmetric.PrivateKey(key),
			Algorithm:  a.(asymmetric.Algorithm),
		}
	}
	return self
}

func (self Keys) AddSymmetricKey(a Algorithm, key []byte) Keys {
	if a.IsSymmetric() {
		self.SymmetricKey = symmetric.Key{
			Secret:    symmetric.Secret(key),
			Algorithm: a.(symmetric.Algorithm),
		}
	}
	return self
}

////////////////////////////////////////////////////////////////////////////////
func (self Keys) GenerateAsymmetricKey(a Algorithm) Keys {
	if a.IsAsymmetric() {
		self.AsymmetricKey = asymmetric.GenerateKeypair(a)
	}
	return self
}

func (self Keys) GenerateSymmetricKey(a Algorithm) Keys {
	if a.IsSymmetric() {
		self.SymmetricKey = symmetric.GenerateKey(a)
	}
}
