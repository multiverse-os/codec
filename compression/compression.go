package compression

import (
	gzip "./gzip"
	none "./none"
	snappy "./snappy"
	zstd "./zstd"
)

type Algorithm int

const (
	None Algorithm = iota
	Gzip
	Snappy
	Zstd
)

// NOTE: Decompress means to relieve the pressure or compression on something,
//       whereas uncompress means to restore a compressed file.
type Compressor interface {
	Compress(input interface{}) ([]byte, error)
	Uncompress(data []byte, output interface{}) error
}

func Load(algorithm Algorithm) Compressor {
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
