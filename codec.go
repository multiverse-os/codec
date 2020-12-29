package codec

type Codec struct {
	format      format.EncodingType
	compression compression.Type
}

func Initialize() Codec {
	return Codec{
		format:      format.Raw,
		compression: compression.None,
	}
}

func (self Codec) FormatType(f format.Type) Codec {
	self.format = format.Load(f)
	return self
}

func (self Codec) CompressionAlgorithm(c compression.Algorithm) Codec {
	self.compression = compression.Load(c)
	return self
}

func (self Codec) Format(input interface{}) ([]byte, error) {
	return self.format.Encode(input)
}

//	Encode(input interface{}) ([]byte, error)
//	Decode(data []byte, output interface{}) error
func (self Codec) Encode(input interface{}) ([]byte, error) {
	return self.Compression.Compress(input)
}

func (self Codec) Decode(data []byte, output interface{}) error {
	return
}
