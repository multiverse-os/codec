package crypto

import (
	asymmetric "github.com/multiverse-os/codec/crypto/key/asymmetric"
	symmetric "github.com/multiverse-os/codec/crypto/key/symmetric"
	params "github.com/multiverse-os/codec/crypto/params"
)

////////////////////////////////////////////////////////////////////////////////
type Key interface {
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)

	Sign(input []byte) ([]byte, error)
	Verify(input []byte) (bool, error)
}

////////////////////////////////////////////////////////////////////////////////
type Cipher struct {
	Key       Key
	Algorithm Algorithm
	Params    params.Params
}

////////////////////////////////////////////////////////////////////////////////
func AsymmetricKey(a Algorithm) Cipher {
	return Cipher{
		Key:       symmetric.Key{},
		Algorithm: a,
		Params:    make(params.Params),
	}
}

func SymmetricKey(a Algorithm) Cipher {
	return Cipher{
		Key:       asymmetric.Keypair{},
		Algorithm: a,
		Params:    make(params.Params),
	}
}

////////////////////////////////////////////////////////////////////////////////

// PARAMS //////////////////////////////////////////////////////////////////////
func (self Cipher) String(name string) string { return self.Params.String(name) }
func (self Cipher) Integer(name string) int   { return self.Params.Integer(name) }
func (self Cipher) Float(name string) float64 { return self.Params.Float(name) }
func (self Cipher) Boolean(name string) bool  { return self.Params.Boolean(name) }

func (self Cipher) AddString(name string, value string) Cipher {
	self.Params.AddString(name, value)
	return self
}

func (self Cipher) AddInteger(name string, value int) Cipher {
	self.Params.AddInteger(name, value)
	return self
}

func (self Cipher) AddFloat(name string, value float64) Cipher {
	self.Params.AddFloat(name, value)
	return self
}

func (self Cipher) AddBoolean(name string, value bool) Cipher {
	self.Params.AddBoolean(name, value)
	return self
}

////////////////////////////////////////////////////////////////////////////////
//func GenerateKey(a algorithm.Type, seed []byte) (key.Key, error) {
//	//return symmetric.Key{
//	//	Algorithm: a,
//	//}, nil
//	//return asymmetric.Keypair{
//	//	Algorithm: a,
//	//}, nil
//	return key.Key{}, nil
//}
