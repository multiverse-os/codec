package sha3

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

type ChecksumData struct{}

func (ChecksumData) Encode(input interface{}) []byte {
	checksum := make([]byte, 64)
	sha3.ShakeSum256(checksum, []byte(fmt.Sprintf("%v", input)))
	return hex.EncodeToString(checksum)

}

func (ChecksumData) String() string { return "sha3" }
