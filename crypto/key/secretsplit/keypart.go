package crypto

import (
	"github.com/multiverse-os/codec"
	"github.com/multiverse-os/codec/checksum"
)

//"As a generalization of the one time pad, SSS, if implemented perfectly, has the unusual property of achieving information theoretic security: it's strong against attackers with unbounded computational power. This kind of "perfect security" adds to the mysterious appeal but this property is of essentially no practical use and seems to distract people from the fact that many attempts at SSS have been so weak a child's speak-n-spell could crack them. Like the one-time-pad and most other cryptosystems, the security of SSS is fairly brittle."
//- Gregory Maxwell [Bitcoin Developer]

type KeyPart []byte

func AssembleKeypair(keyParts ...KeyPart) Keypair {
	var privateKey []byte
	for _, keyPart := range keyParts {
		privateKey = append(privateKey, keyPart)
	}

	keypair := Keypair{
		PrivateKey: privateKey,
		Salt:       codec.Checksum(checksum.XXH64, privateKey),
	}
	keypair.PublicKey = keypair.GeneratePublicKey
	return keypair
}
