package codec

import (
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
	compression compression.Algorithm
}

// Initialization with raw/no compression, that can be defined
func Initialize() Codec {
	return Codec{
		encoding:    encoding.Format(encoding.JSON),
		compression: compression.Type(compression.None),
	}
}

// Chain-able codec defintion //////////////////////////////////////////////////
func (self Codec) FormatType(encodingType encoding.Type) Codec {
	self.encoding = encoding.Format(encodingType)
	return self
}

func (self Codec) CompressionAlgorithm(c compression.CompressionType) Codec {
	self.compression = compression.Type(c)
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
