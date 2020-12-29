package gob

import (
	"bytes"
	"encoding/gob"
)

type DataFormat struct{}

func (DataFormat) Encode(obj interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(value)
}

func (DataFormat) String() string { return "gob" }
