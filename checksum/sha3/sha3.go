package sha3

import (
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"
)

type ChecksumData struct{}

func (ChecksumData) Encode(input interface{}) []byte {
	checksum := make([]byte, 64)
	sha3.ShakeSum256(checksum, []byte(fmt.Sprintf("%v", input)))
	return checksum
}

func (self ChecksumData) Validate(checksum []byte, input []byte) bool {
	return reflect.DeepEqual(checksum, self.Encode(input))
}

func (ChecksumData) String() string { return "sha3" }
