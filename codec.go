package codec

import (
	"fmt"

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
}

// Initialization with raw/no compression, that can be defined
func Initialize() Codec {
	return Codec{
		encoding:    encoding.Format(encoding.Raw),
		compression: compression.Algorithm(compression.None),
	}
}

func (self Codec) String() string {
	return fmt.Sprintf("ecoding=%s,compression=%s", self.encoding.String(), self.compression.String())
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
