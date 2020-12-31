package main

import (
	"fmt"

	xxhash "github.com/multiverse-os/starshipyard/framework/datastore/leveldb/codec/checksum/xxhash"
)

func main() {
	fmt.Println("xxhash exmaple")

	csum := xxhash.ChecksumFormat{}

	output, err := csum.Encode("test11")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("output of test checksum:", output)

}
