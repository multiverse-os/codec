package format

import (
	"bytes"
	"encoding/gob"
)

type gobFormat struct{}

func (gobFormat) Encode(obj interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (gobFormat) Decode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(value)
}

func (gobFormat) String() string { return "gob" }
