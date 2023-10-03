package codec

import (
	"fmt"
	"math/rand"
	"time"

	checksum "github.com/multiverse-os/codec/checksum"
	compression "github.com/multiverse-os/codec/compression"
	encoding "github.com/multiverse-os/codec/encoding"
)

// MarshalUnmarshaler represents a codec used to marshal and unmarshal entities.
type MarshalUnmarshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(b []byte, v interface{}) error
	// name of this codec
	Name() string
}

////////////////////////////////////////////////////////////////////////////////

type Codec struct {
	encoding    encoding.Codec
	compression compression.Codec
	checksum    checksum.Hash
	//keyring     crypto.Keyring
}

// Initialization with raw/no compression, that can be defined
func Initialize() Codec {
	rand.Seed(time.Now().UTC().UnixNano())
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(compression.None),
		checksum:    checksum.Algorithm(checksum.None),
		//keyring:     crypto.EmptyKeyring(),
	}
}

func (self Codec) String() string {
	return fmt.Sprintf("ecoding=%s,compression=%s,checksum=%s", self.encoding.String(), self.compression.String(), self.checksum.String())
	//return fmt.Sprintf("ecoding=%s,compression=%s,checksum=%s,cryptography=%s", self.encoding.String(), self.compression.String(), self.checksum.String(), self.cryptography.Algorithm.String())
}

// Initialize Codec ////////////////////////////////////////////////////////////
func EncodingFormat(encodingType encoding.Type) Codec {
	return Codec{
		encoding:    encoding.Format(encodingType),
		compression: compression.Algorithm(compression.None),
		checksum:    checksum.Algorithm(checksum.None),
		//keyring:     crypto.EmptyKeyring(),
	}
}

func CompressionAlgorithm(c compression.Type) Codec {
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(c),
		checksum:    checksum.Algorithm(checksum.None),
		//keyring:     crypto.EmptyKeyring(),
	}
}

func ChecksumHash(c checksum.Type) Codec {
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(compression.None),
		checksum:    checksum.Algorithm(c),
		//keyring:     crypto.EmptyKeyring(),
	}
}

//func CryptoKey(a crypto.Algorithm) Codec {
//	return Codec{
//		encoding:    encoding.Format(encoding.Raw),
//		compression: compression.Algorithm(compression.None),
//		checksum:    checksum.Algorithm(checksum.None),
//		keyring:     crypto.EmptyKeyring(),
//	}
//}

// Chain-able codec defintion //////////////////////////////////////////////////
func (self Codec) EncodingFormat(encodingType encoding.Type) Codec {
	self.encoding = encoding.Format(encodingType)
	return self
}

func (self Codec) CompressionAlgorithm(c compression.Type) Codec {
	self.compression = compression.Algorithm(c)
	return self
}

func (self Codec) ChecksumAlgorithm(c checksum.Type) Codec {
	self.checksum = checksum.Algorithm(c)
	return self
}

//func (self Codec) CryptoKey(a crypto.Algorithm) Codec {
//	self.cryptography = crypto.CryptoKey(a)
//	return self
//}

// Encode/Decode Format ////////////////////////////////////////////////////////
//
//	Encode(input interface{}) ([]byte, error)
func (self Codec) Encode(input interface{}) ([]byte, error) {
	return self.encoding.Encode(input)
}

// Decode(data []byte, output interface{}) error
func (self Codec) Decode(input []byte, output interface{}) error {
	return self.encoding.Decode(input, output)
}

// Compress / Uncompress ///////////////////////////////////////////////////////
func (self Codec) Compress(input []byte) ([]byte, error) {
	return self.compression.Compress(input)
}

func (self Codec) Uncompress(input []byte) ([]byte, error) {
	return self.compression.Uncompress(input)
}

// Cryptography Functions //////////////////////////////////////////////////////
//func (self Codec) Encrypt(input []byte) ([]byte, error) {
//	return self.cryptography.Key.Encrypt(input)
//}
//
//func (self Codec) Decrypt(input []byte) ([]byte, error) {
//	return self.cryptography.Key.Decrypt(input)
//}
//
//func (self Codec) Sign(input []byte) ([]byte, error) {
//	return self.cryptography.Key.Sign(input)
//}
//
//func (self Codec) VerifySignature(input []byte) (bool, error) {
//	return self.cryptography.Key.Verify(input)
//}

////////////////////////////////////////////////////////////////////////////////

// Checksum ////////////////////////////////////////////////////////////////////

func (self Codec) Checksum(input []byte) []byte {
	return self.checksum.Encode(input)
}

// TODO Add a validation checking.

// MAYBE Add ability to take a large file, and break
// it up into chunks and essentially merkle checksum it.

// //////////////////////////////////////////////////////////////////////////////
// No Codec Access /////////////////////////////////////////////////////////////
func Encode(e encoding.Type, input interface{}) ([]byte, error) {
	return encoding.Format(e).Encode(input)
}

func Decode(e encoding.Type, input []byte, value interface{}) error {
	return encoding.Format(e).Decode(input, value)
}

// ----------------------------------------------------------------------------//
func Compress(c compression.Type, input []byte) ([]byte, error) {
	return compression.Algorithm(c).Compress(input)
}

func Uncompress(c compression.Type, input []byte) ([]byte, error) {
	return compression.Algorithm(c).Uncompress(input)
}

// ----------------------------------------------------------------------------//
func Checksum(c checksum.Type, input []byte) []byte {
	return checksum.Algorithm(c).Encode(input)
}

//----------------------------------------------------------------------------//
//func Encrypt(a crypto.Algorithm, key, input []byte) ([]byte, error) {
//	return crypto.CryptoKey(a).Encrypt(key, input)
//}
//
//func Decrypt(a crypto.Algorithm, key, input []byte) ([]byte, error) {
//	return crypto.Cipher(a).Decrypt(key, input)
//}
//
//func Sign(a crypto.Algorithm, privateKey, input []byte) ([]byte, error) {
//	return crypto.Cipher(a).Sign(privateKey, input)
//}
//
//func VerifySignature(a crypto.Algorithm, publicKey, input []byte) (bool, error) {
//	return crypto.Cipher(a).VerifySignature(publicKey, input)
//}

////////////////////////////////////////////////////////////////////////////////
