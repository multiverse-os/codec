package compression

import (
	gzip "github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/gzip"
	none "github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/none"
	snappy "github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/snappy"
	zstd "github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/zstd"
)

type CompressionType int

const (
	None CompressionType = iota
	Gzip
	Snappy
	Zstd
)

// NOTE: Decompress means to relieve the pressure or compression on something,
//       whereas uncompress means to restore a compressed file.
type Algorithm interface {
	Compress(input []byte) ([]byte, error)
	Uncompress(data []byte) ([]byte, error)
}

func Type(algorithm CompressionType) Algorithm {
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
