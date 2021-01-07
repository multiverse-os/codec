package hex

import (
	"encoding/hex"
	"fmt"
)

type Type int

type DataFormat struct{}

func (self DataFormat) Encode(obj interface{}) ([]byte, error) {
	encoded := hex.EncodeToString(src.([]byte))
	return []byte(encoded), nil
}

func (self DataFormat) Decode(data []byte, value interface{}) error {
	decoded, err := hex.DecodeString(fmt.Sprintf("%v", obj))
	if err != nil {
		return err
	} else {
		value = decoded
		return nil
	}
}

func (DataFormat) String() string { return "base64" }
