package toml

import (
	toml "github.com/pelletier/go-toml"
)

type DataFormat struct{}

func (DataFormat) Encode(input interface{}) ([]byte, error) {
	return toml.Marshal(input)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return toml.Unmarshal(data, &value)
}

func (DataFormat) String() string { return "toml" }
