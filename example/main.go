package main

import (
	"fmt"

	"github.com/multiverse-os/starshipyard/framework/datastore/codec"
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/compression"
	"github.com/multiverse-os/starshipyard/framework/datastore/codec/encoding"
)

type TestObject struct {
	ID          string
	Name        string
	Description string
}

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
	fmt.Println("[TEST][codec]=>[encoding][NONE] -> .Encode(\"test\") ")

	fmt.Println("\n[codec][status]:", codec.String())
	testString := []byte("test")

	fmt.Printf("\n[testString]:%s\n", testString)
	noEncodingBytes, err := codec.Encode(testString)
	if err != nil {
		fmt.Println("[error]:", err)
	}
	fmt.Println("[encoding][NONE][byte]:", noEncodingBytes)
	fmt.Printf("[encoding][NONE][string]:%s\n", noEncodingBytes)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[TEST][codec]=>[compression][NONE]-> .Compress([]byte(\"test\")) ")

	testString = []byte("test")
	fmt.Printf("[testString]:%s\n", test)
	noCompressionBytes, err := codec.Compress([]byte("test"))
	if err != nil {
		fmt.Println("[error]:", err)
	}
	fmt.Println("[compression][NONE][byte]:", noCompressionBytes)
	fmt.Printf("[compression][NONE][string]:%s\n", noCompressionBytes)

	fmt.Println("======================================================")
	fmt.Println("(2) Run through json encoding with no compression")

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[TEST][codec]=[JSON]-> .Encode(\"test\") ")

	codec = codec.EncodingFormat(encoding.JSON)
	testObject := TestObjectFactory()
	emptyObject := &TestObject{}

	fmt.Println("\n[codec][status]:", codec.String())

	fmt.Printf("\n[testObject]:%s\n", testObject)
	encodedObject, err := codec.Encode(testObject)
	if err != nil {
		fmt.Println("[error]:", err)
	}
	fmt.Println("[encodedObject][encoded]byte]:", encodedObject)
	fmt.Printf("[encodedObject][encoded][string]:%s\n", encodedObject)

	fmt.Println("-------------------------------------------------------")

	fmt.Printf("[testObject]:%s\n", testObject)
	err = codec.Decode(encodedObject, &emptyObject)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("[emptyObject][decoded][byte]:", emptyObject)
	fmt.Printf("[emptyObject][decoded][string]:%s\n", emptyObject)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[Length of encoded object]:", len(encodedObject))

	fmt.Println("======================================================")
	fmt.Println("(3) Run through json encoding with gzip compression")

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[TEST][codec]=[encoding][JSON]=[compression][GZIP]-> .Encode(\"test\") ")

	codec = codec.EncodingFormat(encoding.JSON).CompressionAlgorithm(compression.Gzip)
	testObject = TestObjectFactory()
	emptyObject = &TestObject{}

	fmt.Println("\n[codec][status]:", codec.String())
	fmt.Println("\n[testObject]:", testObject)

	encodedObject, err = codec.Encode(testObject)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("[encodedObject][encoded][byte]:", encodedObject)
	fmt.Printf("[encodedObject][encoded][string]:%s\n", encodedObject)

	fmt.Println("[Length of encoded object]:", len(encodedObject))

	fmt.Println("-------------------------------------------------------")
	compressedObject, err := codec.Compress(encodedObject)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("[testObject][compressed][byte]:", compressedObject)
	fmt.Printf("[testObject][compressed][string]:%s\n", compressedObject)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("[Length of encoded object]:", len(encodedObject))
	fmt.Println("[Length of compressed object]:", len(compressedObject))

	fmt.Println("======================================================")
	fmt.Println("(4) Run through zstd compression with raw encoding")

	fmt.Println("[TEST][codec]=[encoding][raw]=[compression][ZSTD]-> .Encode(\"test\") ")
	fmt.Println("-------------------------------------------------------")

	codec = codec.EncodingFormat(encoding.Raw).CompressionAlgorithm(compression.Zstd)
	testObject = TestObjectFactory()

	emptyObject = &TestObject{}

	fmt.Println("[ZSTD]")
	compressedObject, err = codec.Compress(encodedObject)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("[ZSTD]")
	fmt.Println("[testObject][compressed][byte]:", compressedObject)
	fmt.Printf("[testObject][compressed][string]:%s\n", compressedObject)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("[Length of compressed object]:", len(compressedObject))

	fmt.Println("======================================================")

}

func TestObjectFactory() TestObject {
	return TestObject{
		ID:          "TestID",
		Name:        "Tester Name",
		Description: "A test object for passing through the codec",
	}
}
