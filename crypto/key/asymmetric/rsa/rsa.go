package rsa

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"math/rand"
	"strings"
)

type PrivateKey rsa.PrivateKey
type PublicKey rsa.PublicKey
type SessionKey []byte

////////////////////////////////////////////////////////////////////////////////
type KeyType int

const (
	Public KeyType = iota
	Private
)

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
type Keypair struct {
	PrivateKey PrivateKey
	PublicKey  PublicKey

	// TODO: Need ability to store information about symmetrically encryping the
	// Private key.

	SessionKey SessionKey

	BlockSize int
	Hash      hash.Hash
}

////////////////////////////////////////////////////////////////////////////////
// TODO: This needs to be in defaultparams and should be stored as an int under
// blocksize
func BlockSizes() []int      { return []int{2048, 4096, 8192} }
func DefaultHash() hash.Hash { return sha512.New() }
func DefaultBlockSize() int  { return 4096 }

func GenerateKeypair() Keypair {
	privateKey, _ := rsa.GenerateKey(cryptorand.Reader, 4096)
	return Keypair{
		PrivateKey: PrivateKey(*privateKey),
		PublicKey:  PublicKey(privateKey.PublicKey),
	}
}

func GenerateDeterministicKeypair(seed []byte) Keypair {
	paddingLength := DefaultBlockSize() - len(seed)
	if paddingLength < 0 {
		seed = seed[:DefaultBlockSize()]
	} else if 0 < paddingLength {
		//padding, _ := GenerateRandomBytes(paddingLength)
		padding := []byte(strings.Repeat("0", paddingLength))
		seed = append(seed, padding...)
	}
	privateKey, _ := rsa.GenerateKey(rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(seed)))), DefaultBlockSize())
	return Keypair{
		PrivateKey: PrivateKey(*privateKey),
		PublicKey:  PublicKey(privateKey.PublicKey),
	}
}

////////////////////////////////////////////////////////////////////////////////
func (self PrivateKey) RSA() *rsa.PrivateKey {
	key := rsa.PrivateKey(self)
	return &(key)
}

func (self PublicKey) RSA() *rsa.PublicKey {
	key := rsa.PublicKey(self)
	return &(key)
}

////////////////////////////////////////////////////////////////////////////////
func (self PrivateKey) Bytes() []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(self.RSA()),
		},
	)
}

func (self PublicKey) Bytes() []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(self.RSA())
	if err != nil {
		panic(err)
	}
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})
}

func (self SessionKey) Bytes() []byte { return []byte(self) }

////////////////////////////////////////////////////////////////////////////////
func LoadPrivateKey(privateKeyBytes []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(privateKeyBytes)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println("error:", err)
	}
	return key
}

func LoadPublicKey(publicKeyBytes []byte) *rsa.PublicKey {
	block, _ := pem.Decode(publicKeyBytes)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		fmt.Println("error:", err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		fmt.Println("error: not ok")
	}
	return key
}

////////////////////////////////////////////////////////////////////////////////
func (self Keypair) Cryptosystem() string {
	return "asymmetric"
}

////////////////////////////////////////////////////////////////////////////////
// TODO: WE are generating a symmetric key; it would be nice to use our
// symmetric key system to tie it altogether nicely.
func (self Keypair) GenerateSessionKey() SessionKey {
	// The hybrid scheme should use at least a 16-byte symmetric key. Here
	// we read the random key that will be used if the RSA decryption isn't
	// well-formed.
	key := make([]byte, 32)

	// TODO: Perhaps instead of just doing a random value, we use our private key
	// and integer index, so we can deterministically create session keys
	// (optionally).

	if _, err := io.ReadFull(cryptorand.Reader, key); err != nil {
		panic("RNG failure")
	}
	// TODO: This would only work if we had Keypair be a pointer
	//self.SessionKey = SessionKey(key)
	return SessionKey(key)

	//   // Given the resulting key, a symmetric scheme can be used to decrypt a
	//   // larger ciphertext.
	//   block, err := aes.NewCipher(key)
	//   if err != nil {
	//       panic("aes.NewCipher failed: " + err.Error())
	//   }
	//
	//   // Since the key is random, using a fixed nonce is acceptable as the
	//   // (key, nonce) pair will still be unique, as required.
	//   var zeroNonce [12]byte
	//   aead, err := cipher.NewGCM(block)
	//   if err != nil {
	//       panic("cipher.NewGCM failed: " + err.Error())
	//   }
	//   ciphertext, _ := hex.DecodeString("00112233445566")
	//   plaintext, err := aead.Open(nil, zeroNonce[:], ciphertext, nil)
	//   if err != nil {
	//       // The RSA ciphertext was badly formed; the decryption will
	//       // fail here because the AES-GCM key will be incorrect.
	//       fmt.Fprintf(os.Stderr, "Error decrypting: %s\n", err)
	//       return
	//   }
	//
	//   fmt.Printf("Plaintext: %s\n", string(plaintext))

}

func (self Keypair) EncryptWithSessionKey(plainText []byte) []byte {
	//rsaCiphertext, _ := hex.DecodeString(plainText)

	if err := rsa.DecryptPKCS1v15SessionKey(cryptorand.Reader, self.PrivateKey.RSA(), plainText, self.SessionKey); err != nil {
		// Any errors that result will be “public” – meaning that they
		// can be determined without any secret information. (For
		// instance, if the length of key is impossible given the RSA
		// public key.)
		//fmt.Fprintf(os.Stderr, "Error from RSA decryption: %s\n", err)
		panic(err)
	}
	return []byte{}
}

////////////////////////////////////////////////////////////////////////////////
func (self Keypair) Encrypt(plainText []byte) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, cryptorand.Reader, self.PublicKey.RSA(), plainText, nil)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func (self Keypair) Decrypt(cipherText []byte) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, cryptorand.Reader, self.PrivateKey.RSA(), cipherText, nil)
	if err != nil {
		panic(err)
	}
	return plaintext
}

////////////////////////////////////////////////////////////////////////////////
func (self Keypair) Sign(input []byte) ([]byte, error) {
	// TODO: Add ability to set options, the option allowed is customizing the
	// hash function which could be nice, if we want to say use xxhash
	return self.PrivateKey.RSA().Sign(cryptorand.Reader, input, nil)
}

func (self Keypair) Verify(input, signature []byte) (bool, error) {
	hashedInput := sha512.Sum512(input)
	signatureHex, _ := hex.DecodeString(string(signature))

	err := rsa.VerifyPKCS1v15(self.PublicKey.RSA(), crypto.SHA256, hashedInput[:], signatureHex)
	if err != nil {
		return false, err
	}
	return true, nil
}
