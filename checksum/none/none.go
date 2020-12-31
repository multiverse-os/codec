package none

import (
	"fmt"
)

type ChecksumData struct{}

func (ChecksumData) Encode(input interface{}) []byte {
	return []byte(fmt.Sprintf("%v", input))
}

func (ChecksumData) String() string { return "none" }
