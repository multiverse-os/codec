package hex

import (
	"encoding/base64"
)

type Type int

type DataFormat struct{}

func (self DataFormat) Encode(obj interface{}) ([]byte, error) {
	var encodedData string
	encodedData = base64.URLEncoding.EncodeToString([]byte(obj.([]byte)))
	return []byte(encodedData), nil
	//var buffer bytes.Buffer
	//encoder := base64.NewEncoder(base64.URLEncoding, &buffer)
	//err := json.NewEncoder(encoder).Encode(obj)
	//if err != nil {
	//	return []byte{}, err
	//}
	//encoder.Close()
	//return buffer.Bytes(), nil
}

func (self DataFormat) Decode(data []byte, value interface{}) error {
	value = base64.URLEncoding.EncodeToString([]byte(data))
	return nil
}

func (DataFormat) String() string { return "base64" }
