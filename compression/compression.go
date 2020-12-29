package compression

import (
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/gzip"
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/none"
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/snappy"
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/zstd"
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

func Load(compressionType Algorithm) Compressor {
	switch compressionType {
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
