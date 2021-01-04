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

	argon2 := crypto.Argon2
	fmt.Println("cryptosystem:", argon2)
	if argon2.IsSymmetric() {
		fmt.Println("IS SYMMETRIC")
	} else {
		fmt.Println("IS ASYMMETRIC")
	}

	rsa := crypto.RSA
	fmt.Println("cryptosystem:", rsa)
	if rsa.IsSymmetric() {
		fmt.Println("IS SYMMETRIC")
	} else {
		fmt.Println("IS ASYMMETRIC")
	}

	cipher := crypto.Cipher(crypto.RSA)

	fmt.Println("Adding data to our cipher that can be used to configure the")
	fmt.Println("the encryption.\n")
	fmt.Println("      cipher.AddString(\"test\", \"test\")\n")
	cipher.AddString("test", "test")

	fmt.Println("Now trying to pull the data out:")
	fmt.Println("      cipher.String(\"test\")")
	fmt.Println("string with name test:", cipher.String("test"))

	argon2Cipher := crypto.Cipher(crypto.Argon2)

	argon2Cipher.Params = argon2.DefaultParams(argon2Cipher.Params)

}
