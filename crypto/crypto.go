package crypto

import (
	"fmt"

	key "github.com/multiverse-os/codec/crypto/key"
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
type cipher struct {
	Key
	Algorithm key.Algorithm
	Params    params.Params
}

////////////////////////////////////////////////////////////////////////////////
func Cipher(a key.Algorithm) cipher {
	var key Key
	if a.IsAsymmetric() {
		fmt.Println("algorithm is asymmetric")
		key = asymmetric.Keypair{}
	} else {
		fmt.Println("algorithm is symmetric")
		key = symmetric.Key{}
	}
	return cipher{
		Algorithm: a,
		Key:       key,
		Params:    make(params.Params),
	}
}

// PARAMS //////////////////////////////////////////////////////////////////////
func (self cipher) String(name string) string { return self.Params.String(name) }
func (self cipher) Integer(name string) int   { return self.Params.Integer(name) }
func (self cipher) Float(name string) float64 { return self.Params.Float(name) }
func (self cipher) Boolean(name string) bool  { return self.Params.Boolean(name) }

func (self cipher) AddString(name string, value string) cipher {
	self.Params.AddString(name, value)
	return self
}

func (self cipher) AddInteger(name string, value int) cipher {
	self.Params.AddInteger(name, value)
	return self
}

func (self cipher) AddFloat(name string, value float64) cipher {
	self.Params.AddFloat(name, value)
	return self
}

func (self cipher) AddBoolean(name string, value bool) cipher {
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
