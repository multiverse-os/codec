package main

import (
	"fmt"

	"github.com/multiverse-os/starshipyard/framework/datastore/codec"
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/encoding"
)

func main() {
	fmt.Println("CODEC EXAMPLE v0.1.0")
	fmt.Println("======================================================")
	fmt.Println("An example program that explains and provides basic testing")
	fmt.Println("together...")

	fmt.Println("[] Initializing the codec object, which stores: ")
	fmt.Println(" '--> An encoding format [bson, json, cbor, etc...]")
	fmt.Println(" '--> A compression algorithm [gzip, zstd, snappy, etc...]\n\n")
	codec := codec.Initialize()

	fmt.Println("The codec object by default is configured to use raw/none")
	fmt.Println("And so has to be configured to do anything but pass the data")
	fmt.Println("right through unchanged.")
	fmt.Println("It is configured using the .Chain().Chain() Style:\n")
	fmt.Println("   codec.Initialize().Encoding(encoding.JSON).Compression(compression.Gzip)\n\n")
	fmt.Println("And it can be reconfigured on the fly, and then the:")
	fmt.Println("    codec.Compress(\"data\")")
	fmt.Println("    [OR] ")
	fmt.Println("    codec.Encode(\"test\")\n")

	fmt.Println("======================================================")
	fmt.Println("(1) Run through raw and no compression")
	var test string
	fmt.Println("-------------------------------------------------------")
	fmt.Println("[TEST] Codec= NO Encoding Type -> .Encode(\"test\") ")
	test = "test"
	fmt.Printf("test:%s\n", test)
	if test, err := codec.Encode("test"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("test:", test)
		fmt.Printf("test:%s\n", test)
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[TEST] Codec= NO Compression Algorithm -> .Compress([]byte(\"test\")) ")

	test = "test"
	fmt.Printf("test:%s\n", test)
	if test, err := codec.Compress([]byte("test")); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("test:", test)
		fmt.Printf("test:%s\n", test)
	}

	fmt.Println("======================================================")
	fmt.Println("(2) Run through json encoding with no compression")

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[TEST] Codec= JSON Encoding Type -> .Encode(\"test\") ")

	codec = codec.EncodingFormat(encoding.JSON)

	fmt.Println("codec encoding format:", codec.String())

	type What struct {
		Test string
		Who  string
	}

	tester := What{
		Test: "test",
		Who:  "What",
	}

	fmt.Printf("test:%s\n", test)
	testo, err := codec.Encode(tester)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("encodo:", testo)
	fmt.Printf("encodo:%s\n", testo)

	returno := &What{}

	if err := codec.Decode(testo, &returno); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("returno:", returno)
		fmt.Printf("returno:%s\n", returno)
	}

	fmt.Println("-------------------------------------------------------")

	fmt.Println("======================================================")
	fmt.Println("(3) Run through json encoding with gzip compression")
	fmt.Println("======================================================")
	fmt.Println("(4) Run through zstd compression with raw encoding")
	fmt.Println("======================================================")
	fmt.Println("(5) Run through and back a few differnet combos...")
	fmt.Println("======================================================")

}
