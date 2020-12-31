package checksum

import (
	none "github.com/multiverse-os/starshipyard/framework/datastore/leveldb/codec/checksum/none"
	sha3 "github.com/multiverse-os/starshipyard/framework/datastore/leveldb/codec/checksum/sha3"
	xxh64 "github.com/multiverse-os/starshipyard/framework/datastore/leveldb/codec/checksum/xxh64"
)

type Hash int

const (
	None Hash = iota
	XXH64
	SHA3
	// ChaCha
	// ..
)

type Checksum interface {
	Encode(input interface{}) ([]byte, error)
	Validate(checksum []byte, data []byte) error
	String() string
}

func Algorithm(dataType Hash) Checksum {
	switch dataType {
	case XXH:
		return xxh64.ChecksumData{}
	case SHA3:
		return sha3.ChecksumData{}
	default: // None
		return none.ChecksumData{}
	}
}
