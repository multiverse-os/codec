package crypto

import (
	asymmetric "github.com/multiverse-os/codec/crypto/key/asymmetric"
	symmetric "github.com/multiverse-os/codec/crypto/key/symmetric"
)

////////////////////////////////////////////////////////////////////////////////
type Keys struct {
	AsymmetricKeys []*asymmetric.Keypair
	SymmetricKeys  []*symmetric.Key
}

////////////////////////////////////////////////////////////////////////////////
func EmptyKeyring() Keys {
	return Keys{
		AsymmetricKeys: []*asymmetric.Keypair{},
		SymmetricKeys:  []*symmetric.Key{},
	}
}

func (self Keys) PublicKey(a Algorithm, key []byte) Keys {
	if a.IsAsymmetric() {
		self.AsymmetricKeys = append(self.AsymmetricKeys, &asymmetric.Keypair{
			PublicKey: asymmetric.PublicKey(key),
			Algorithm: a.(asymmetric.Algorithm),
		})
	}
	return self
}

func (self Keys) PrivateKey(a Algorithm, key []byte) Keys {
	if a.IsAsymmetric() {
		self.AsymmetricKeys = append(self.AsymmetricKeys, &asymmetric.Keypair{
			PrivateKey: asymmetric.PrivateKey(key),
			Algorithm:  a.(asymmetric.Algorithm),
		})
	}
	return self
}

func (self Keys) SymmetricKey(a Algorithm, key []byte) Keys {
	if a.IsSymmetric() {
		self.SymmetricKeys = append(self.SymmetricKeys, &symmetric.Key{
			Secret:    symmetric.Secret(key),
			Algorithm: a.(symmetric.Algorithm),
		})
	}
	return self
}
