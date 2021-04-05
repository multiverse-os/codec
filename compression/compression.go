package compression

import (
	gzip "github.com/multiverse-os/codec/compression/gzip"
	none "github.com/multiverse-os/codec/compression/none"
	snappy "github.com/multiverse-os/codec/compression/snappy"
	zstd "github.com/multiverse-os/codec/compression/zstd"
)

////////////////////////////////////////////////////////////////////////////////
var (
	magicZIP  = []byte{0x50, 0x4b, 0x03, 0x04}
	magicGZ   = []byte{0x1f, 0x8b}
	magicBZIP = []byte{0x42, 0x5a}
	magicTAR  = []byte{0x75, 0x73, 0x74, 0x61, 0x72} // at offset 257
	magicXZ   = []byte{0xfd, 0x37, 0x7a, 0x58, 0x5a, 0x00}
)

////////////////////////////////////////////////////////////////////////////////
type Type int

const (
	None Type = iota
	Gzip
	Snappy
	Zstd
)

// NOTE: Decompress means to relieve the pressure or compression on something,
//       whereas uncompress means to restore a compressed file.
type Codec interface {
	Compress(input []byte) ([]byte, error)
	Uncompress(data []byte) ([]byte, error)
	String() string
}

func Algorithm(algorithm Type) Codec {
	switch algorithm {
	case Gzip:
		return gzip.Compression{}
	case Snappy:
		return snappy.Compression{}
	case Zstd:
		return zstd.Compression{}
	default: // None
		return none.Compression{}
	}

}
