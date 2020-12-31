package raw

import (
	"bytes"
	"encoding/binary"
)

type DataFormat struct{}

func (DataFormat) Encode(input interface{}) ([]byte, error) {
	//b := *(*[unsafe.Sizeof(value)]byte)(unsafe.Pointer(&value))
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, input)
	return buffer.Bytes(), nil
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	binary.Read(bytes.NewBuffer(data[:]), binary.LittleEndian, &value)
	return nil
}

func (DataFormat) String() string { return "raw" }
