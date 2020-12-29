package zstd

import (
	zstd "github.com/klauspost/compress/zstd"
)

// TODO: Eventually options will go here
type Compression struct{}

func (Compression) Compress(input []byte) ([]byte, error) {
	encoder, _ := zstd.NewWriter(nil)
	return encoder.EncodeAll(input, make([]byte, 0, len(input))), nil
}

func (Compression) Uncompress(input []byte) ([]byte, error) {
	var output []byte
	decoder, _ := zstd.NewReader(nil)
	decoder.DecodeAll(input, output)
	return output, nil
}

func (Compression) String() string {
	return "zstd"
}
