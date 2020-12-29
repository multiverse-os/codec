package cbor

type DataFormat struct{}

func (DataFormat) Encode(obj interface{}) ([]byte, error) {
	//return cbor.Marshal(obj)
	return []byte{}, nil
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	//return cbor.Unmarshal(data, &value)
	return nil
}

func (DataFormat) String() string { return "cbor" }
