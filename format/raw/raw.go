package codec

type DataFormat struct{}

func (DataFormat) Encode(obj interface{}) ([]byte, error) {
	return obj
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return data
}

func (DataFormat) String() string { return "raw" }
