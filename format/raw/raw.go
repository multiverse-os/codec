package codec

type rawFormat struct{}

func (rawFormat) Encode(obj interface{}) ([]byte, error) {
	return obj
}

func (rawFormat) Decode(data []byte, value interface{}) error {
	return data
}

func (rawFormat) String() string { return "raw" }
