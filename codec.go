package codec

import (
	compression "./compression"
	format "./format"
)

type Codec struct {
	format      format.EncodingType
	compression compression.Type
}

// Initialization with raw/no compression, that can be defined
func Initialize() Codec {
	return Codec{
		format:      format.Raw,
		compression: compression.None,
	}
}

// Chain-able codec defintion //////////////////////////////////////////////////
func (self Codec) FormatType(f format.Type) Codec {
	self.format = format.Load(f)
	return self
}

func (self Codec) CompressionAlgorithm(c compression.Algorithm) Codec {
	self.compression = compression.Load(c)
	return self
}

// Encode/Decode Format ////////////////////////////////////////////////////////
//	Encode(input interface{}) ([]byte, error)
func (self Codec) Encode(input interface{}) ([]byte, error) {
	return self.format.Encode(input)
}

//	Decode(data []byte, output interface{}) error
func (self Codec) Decode(input interface{}, output interface{}) error {
	return self.format.Decode(input, output)
}

// Compress / Uncompress ///////////////////////////////////////////////////////
func (self Codec) Compress(input interface{}) ([]byte, error) {
	return self.compression.Compress(input)
}

func (self Codec) Uncompress(input interface{}, output interface{}) error {
	return self.compression.Uncompress(input)
}
