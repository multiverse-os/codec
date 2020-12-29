package main

import (
	"fmt"

	"github.com/multiverse-os/starshipyard/framework/datastore/codec"
)

func main() {
	fmt.Println("codec example")
	fmt.Println("======================================================")
	fmt.Println("(1) Run through raw and no compression")

	codec := codec.Initialize()

	if test, err := codec.Encode("test"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("test:", test)
		fmt.Printf("test:%s\n", test)
	}

	fmt.Println("======================================================")
	fmt.Println("(2) Run through json encoding with no compression")
	fmt.Println("======================================================")
	fmt.Println("(3) Run through json encoding with gzip compression")
	fmt.Println("======================================================")
	fmt.Println("(4) Run through zstd compression with raw encoding")
	fmt.Println("======================================================")
	fmt.Println("(5) Run through and back a few differnet combos...")
	fmt.Println("======================================================")

}
