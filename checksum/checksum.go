package checksum

import (
	none "github.com/multiverse-os/levelup/data/codec/checksum/none"
	sha3 "github.com/multiverse-os/levelup/data/codec/checksum/sha3"
	xxh64 "github.com/multiverse-os/levelup/data/codec/checksum/xxh64"
)

type Type int

const (
	None Type = iota
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
	case XXH64:
		return xxh64.ChecksumData{}
	case SHA3:
		return sha3.ChecksumData{}
	default: // None
		return none.ChecksumData{}
	}
}
