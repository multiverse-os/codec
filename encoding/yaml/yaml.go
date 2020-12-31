package yaml

import (
	yaml "gopkg.in/yaml.v2"
)

type DataFormat struct{}

func (DataFormat) Encode(input interface{}) ([]byte, error) {
	return yaml.Marshal(&input)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return yaml.Unmarshal(data, &value)
}

func (DataFormat) String() string { return "yaml" }
