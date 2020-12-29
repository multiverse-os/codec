package json

import (
	"encoding/json"
	"fmt"
)

type DataFormat struct{}

func (DataFormat) Encode(obj interface{}) ([]byte, error) {
	fmt.Println("json marshalling")
	return json.Marshal(obj)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return json.Unmarshal(data, &value)
}

func (DataFormat) String() string { return "json" }
