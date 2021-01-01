package main

import (
	"fmt"

	crypto "github.com/multiverse-os/codec/crypto"
)

func main() {
	fmt.Println("metakey: cryptography key library")
	fmt.Println("===============================================")
	fmt.Println("an advanced cryptogrpahic library that provides a wide range")
	fmt.Println("of functionality for algorithms that do not commonly support it")

	keypair, err := crypto.GenerateKeypair(algorithm.Argon2, []byte("test"))
	if err != nil {
		fmt.Println("[error] what the error:", err)
	}

	fmt.Println("keypair.Algorithm.String():", keypair.Algorithm.String())
	fmt.Println("keypair.PrivateKey():", keypair.PrivateKey())
	fmt.Println("keypair.PublicKey():", keypair.PublicKey())

}
