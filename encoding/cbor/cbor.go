package cbor

type DataFormat struct{}

func (DataFormat) Encode(value interface{}) ([]byte, error) {
	//return Marshal(value)
	return []byte{}, nil
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	//return cbor.Unmarshal(data, &value)
	//cborData, _ := hex.DecodeString("a46341676504644e616d656543616e6479664f776e65727382644d617279634a6f65644d616c65f4")
	//return Unmarshal(data, &value)
	return nil
}

func (DataFormat) String() string { return "cbor" }
