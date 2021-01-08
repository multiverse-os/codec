package xxh32

import (
	"fmt"
	"reflect"
)

type ChecksumData struct{}

func (ChecksumData) Encode(input interface{}) []byte {
	//var buffer bytes.Buffer
	//binary.Write(&buffer, binary.BigEndian, input)
	//d.Write(buffer.Bytes())

	d := New()
	d.Write([]byte(fmt.Sprintf("%v", input)))
	output := d.Sum32()
	return []byte(fmt.Sprintf("%v", output))
}

func (self ChecksumData) Validate(checksum []byte, input []byte) bool {
	return reflect.DeepEqual(checksum, self.Encode(input))
}

func (ChecksumData) String() string { return "xxhash" }
