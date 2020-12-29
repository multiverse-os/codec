package codec

import (
	"encoding/json"
)

type jsonFormat struct{}

func (jsonFormat) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (jsonFormat) Decode(data []byte, value interface{}) error {
	return json.Unmarshal(data, &value)
}

func (jsonFormat) String() string { return "json" }
