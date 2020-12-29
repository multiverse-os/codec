package snappy

import (
	snappy "github.com/klauspost/compress/snappy"
)

type Compression struct{}

func (Compression) Compress(input []byte) ([]byte, error) {
	return snappy.Encode(nil, input), nil
}

func (Compression) Uncompress(input []byte) ([]byte, error) {
	return snappy.Decode(nil, input)
}

func (Compression) String() string {
	return "snappy"
}
