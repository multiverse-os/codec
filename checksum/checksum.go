package checksum

import (
	none "github.com/multiverse-os/codec/checksum/none"
	sha3 "github.com/multiverse-os/codec/checksum/sha3"
	xxh64 "github.com/multiverse-os/codec/checksum/xxh64"
)

type Type int

const (
	None Type = iota
	XXH32
	XXH64
	SHA3
	// ChaCha
	// ..
)

type Hash interface {
	Encode(input interface{}) []byte
	Validate(checksum []byte, data []byte) bool
	String() string
}

func Algorithm(dataType Type) Hash {
	switch dataType {
	case XXH32:
		return xxh32.ChecksumData{}
	case XXH64:
		return xxh64.ChecksumData{}
	case SHA3:
		return sha3.ChecksumData{}
	default: // None
		return none.ChecksumData{}
	}
}
