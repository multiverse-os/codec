package key

import (
	asymmetric "github.com/multiverse-os/codec/crypto/key/asymmetric"
	symmetric "github.com/multiverse-os/codec/crypto/key/symmetric"
)

type Cryptosystem int

func (self Cryptosystem) String() string { return self.String() }

//for _, a := range Algorithms() {
//	if len(self.String()) == len(a.String()) &&
//		self.String() == a.String() {
//		a.String()
//	}
//}
//return ""

func (self Cryptosystem) IsAsymmetric() bool {
	for _, cs := range asymmetric.Algorithms() {
		if len(self.String()) == len(cs.String()) &&
			self.String() == cs.String() {
			return true
		}
	}
	return false
}

func (self Cryptosystem) IsSymmetric() bool {
	for _, cs := range symmetric.Algorithms() {
		if len(self.String()) == len(cs.String()) &&
			self.String() == cs.String() {
			return true
		}
	}
	return false
}
