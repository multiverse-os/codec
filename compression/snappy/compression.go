package snappy

import (
	snappy "github.com/klauspost/compress/snappy"
)

type Compression struct{}

func (Compression) Compress(input []byte) []byte {
	return snappy.Encode(nil, input)
}

func (Compression) Uncompress(input []byte) ([]byte, error) {
	return snappy.Decode(nil, input)
}

func (Compression) String() string {
	return "snappy"
}
