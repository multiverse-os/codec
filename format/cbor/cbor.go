package cbor

type dataFormat struct{}

func (dataFormat) Encode(obj interface{}) ([]byte, error) {
	//return cbor.Marshal(obj)
	return []byte{}, nil
}

func (dataFormat) Decode(data []byte, value interface{}) error {
	//return cbor.Unmarshal(data, &value)
	return nil
}

func (dataFormat) String() string { return "cbor" }
