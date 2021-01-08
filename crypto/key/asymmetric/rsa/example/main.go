package main

import (
	"fmt"

	rsa "github.com/multiverse-os/codec/crypto/key/asymmetric/rsa"
)

func main() {
	keypair := rsa.GenerateKeypair()

	fmt.Println("private key:", keypair.PrivateKey)
	fmt.Println("public key:", keypair.PublicKey)

	dKeypair := rsa.GenerateDeterministicKeypair([]byte("test"))

	fmt.Println("private key:", dKeypair.PrivateKey)
	fmt.Println("public key:", dKeypair.PublicKey)

	fmt.Println("pubPEM:", dKeypair.PublicKey.Bytes())

}
