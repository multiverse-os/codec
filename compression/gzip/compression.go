package gzip

import (
	"bytes"
	"io/ioutil"

	gzip "github.com/multiverse-os/starshipyard/framework/datastore/codec/compression/gzip/gzip"
)

// NOTE: Parallelization of gzip boosts the preformance by:
// 1573.23MB/s 	101.9x 	4902285651 	+2.53%

// https://github.com/klauspost/pgzip/

// GZIP.SetConcurrency(BlockSize, Blocks)
type Compression struct {
	ChunkSize int
	Parallel  int
}

func Load() Gzip {
	return Gzip{
		//ChunkSize: 512000, // 5 Megabytes
		ChunkSize: 102400, // 1 Megabyte
		Parallel:  4,      // 4 Parallel Processes
	}
}

func (self Gzip) Compress(input []byte) []byte {
	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)
	gzipWriter.SetConcurrency(self.ChunkSize, self.Parallel)
	gzipWriter.Write(input)
	gzipWriter.Close()
	return buffer.Bytes()
}

func (Gzip) Uncompress(input []byte) ([]byte, error) {
	gzipReader, _ := gzip.NewReader(bytes.NewReader(input))
	return ioutil.ReadAll(gzipReader)
}

func (Gzip) String() string {
	return "gzip"
}
