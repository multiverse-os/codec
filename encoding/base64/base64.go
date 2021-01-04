package json

import (
	"encoding/base64"
	"encoding/json"
)

type DataFormat struct {
}

func (self DataFormat) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)

	salt_b64 := base64.URLEncoding.EncodeToString(salt)
	pwhash_b64 := base64.URLEncoding.EncodeToString(dk)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return json.Unmarshal(data, &value)
}

func (DataFormat) String() string { return "json" }
