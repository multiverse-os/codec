package cbor

import (
	cbor "./cbor"
)

type DataFormat struct{}

func (DataFormat) Encode(value interface{}) ([]byte, error) {
	return cbor.Marshal(value)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	//return cbor.Unmarshal(data, &value)
	//cborData, _ := hex.DecodeString("a46341676504644e616d656543616e6479664f776e65727382644d617279634a6f65644d616c65f4")
	return cbor.Unmarshal(data, &value)
}

func (DataFormat) String() string { return "cbor" }
