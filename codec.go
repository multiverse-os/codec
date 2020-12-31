package codec

import (
	"fmt"

	checksum "github.com/multiverse-os/starshipyard/framework/datastore/codec/checksum"
	compression "github.com/multiverse-os/starshipyard/framework/datastore/codec/compression"
	encoding "github.com/multiverse-os/starshipyard/framework/datastore/codec/encoding"
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
}

// Initialization with raw/no compression, that can be defined
func Initialize() Codec {
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(compression.None),
		checksum:    checksum.Algorithm(checksum.None),
	}
}

func (self Codec) String() string {
	return fmt.Sprintf("ecoding=%s,compression=%s", self.encoding.String(), self.compression.String())
}

// Initialize Codec ////////////////////////////////////////////////////////////
func EncodingFormat(encodingType encoding.Type) Codec {
	return Codec{
		encoding:    encoding.Format(encodingType),
		compression: compression.Algorithm(compression.None),
		checksum:    checksum.Algorithm(checksum.None),
	}
}

func CompressionAlgorithm(c compression.Type) Codec {
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(c),
		checksum:    checksum.Algorithm(checksum.None),
	}
}

func ChecksumHash(c checksum.Type) Codec {
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(compression.None),
		checksum:    checksum.Algorithm(c),
	}
}

// Chain-able codec defintion //////////////////////////////////////////////////
func (self Codec) EncodingFormat(encodingType encoding.Type) Codec {
	self.encoding = encoding.Format(encodingType)
	return self
}

func (self Codec) CompressionAlgorithm(c compression.Type) Codec {
	self.compression = compression.Algorithm(c)
	return self
}

func (self Codec) ChecksumHash(c checksum.Type) Codec {
	self.checksum = checksum.Algorithm(c)
	return self
}

// Encode/Decode Format ////////////////////////////////////////////////////////
//	Encode(input interface{}) ([]byte, error)
func (self Codec) Encode(input interface{}) ([]byte, error) {
	return self.encoding.Encode(input)
}

//	Decode(data []byte, output interface{}) error
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

////////////////////////////////////////////////////////////////////////////////

// Checksum ////////////////////////////////////////////////////////////////////

func (self Codec) Checksum(input []byte) []byte {
	return self.checksum.Encode(input)
}

// TODO Add a validation checking.

// MAYBE Add ability to take a large file, and break
// it up into chunks and essentially merkle checksum it.

////////////////////////////////////////////////////////////////////////////////
// No Codec Access /////////////////////////////////////////////////////////////

func Encode(e encoding.Type, input interface{}) ([]byte, error) {
	codec := encoding.Format(e)
	return codec.Encode(input)
}

func Decode(e encoding.Type, input []byte, value interface{}) error {
	codec := encoding.Format(e)
	return codec.Decode(input, value)
}

func Compress(c compression.Type, input []byte) ([]byte, error) {
	codec := compression.Algorithm(c)
	return codec.Compress(input)
}

func Uncompress(c compression.Type, input []byte) ([]byte, error) {
	codec := compression.Algorithm(c)
	return codec.Uncompress(input)
}

func Checksum(c checksum.Type, input []byte) []byte {
	hash := checksum.Algorithm(c)
	return hash.Encode(input)
}
