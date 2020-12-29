package codec

import (
	"encoding/json"
)

type DataFormat struct{}

func (DataFormat) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return json.Unmarshal(data, &value)
}

func (DataFormat) String() string { return "json" }
