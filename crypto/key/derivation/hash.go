package keyderivation

import (
	"hash"

	// Hash
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	// Crypto
	"golang.org/x/crypto/md4"
)

var str2hash = map[string](func() hash.Hash){
	"md4":    md4.New,
	"md5":    md5.New,
	"sha1":   sha1.New,
	"sha224": sha256.New224,
	"sha256": sha256.New,
	"sha384": sha512.New384,
	"sha512": sha512.New,
}

//func Valid() (bool, error) {
//	if bcrypt.MinCost < params.Integer("cost") && params.Integer("cost") < bcrypt.MaxCost {
//		return false, fmt.Errorf("Error: ", "bcrypt: unsupported cost value")
//	} else {
//		return true, nil
//	}
//}
//
//func DefaultParams(params params.Params) params.Params {
//
//	//opts.Rounds = 50000
//	//opts.Hashname = "sha256"
//	//opts.Kdname = "scrypt"
//	//opts.Cost = 14
//	//parser := flags
//
//	params.AddInteger("iterations", 5000)
//	params.AddString("hash", "sha256")
//	params.AddString("keyderivation", "name")
//	params.AddInteger("cost", 10)
//	params.AddString("hmac", "base64")
//	return params
//}
