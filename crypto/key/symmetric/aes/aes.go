package aes

import (
	"github.com/multiverse-os/codec/compression/gzip"
	"github.com/multiverse-os/codec/compression/none"
	"github.com/multiverse-os/codec/compression/snappy"
	"github.com/multiverse-os/codec/compression/zstd"
)

////////////////////////////////////////////////////////////////////////////////
type Algorithm int

const (
	None Type = iota
	Gzip
	Snappy
	Zstd
)

////////////////////////////////////////////////////////////////////////////////
// NOTE: Decompress means to relieve the pressure or compression on something,
//       whereas uncompress means to restore a compressed file.
type Codec interface {
	Compress(input []byte) ([]byte, error)
	Uncompress(data []byte) ([]byte, error)
	String() string
}

////////////////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////////////////
type Key struct {
	Secret []byte
}

////////////////////////////////////////////////////////////////////////////////
func (self Key) SecretKey(secretKey []byte) Key {
	self.Secret = secretKey
	return self
}

////////////////////////////////////////////////////////////////////////////////
func (self Key) Cryptosystem() string {
	return "symmetric"
}

func (self Key) Encrypt(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Key) Decrypt(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Key) Sign(input []byte) ([]byte, error) {
	return []byte{}, nil
}

func (self Key) Verify(input []byte) (bool, error) {
	return false, nil
}

////////////////////////////////////////////////////////////////////////////////
