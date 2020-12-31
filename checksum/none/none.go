package none

import (
	"fmt"
	"reflect"
)

type ChecksumData struct{}

func (ChecksumData) Encode(input interface{}) []byte {
	return []byte(fmt.Sprintf("%v", input))
}

func (self ChecksumData) Validate(checksum []byte, input []byte) bool {
	return reflect.DeepEqual(checksum, self.Encode(input))
}

func (ChecksumData) String() string { return "none" }
