package key

import (
	asymmetric "github.com/multiverse-os/codec/crypto/key/asymmetric"
	symmetric "github.com/multiverse-os/codec/crypto/key/symmetric"
)

type Algorithm interface {
	String() string

	IsAsymmetric() bool
	IsSymmetric() bool
}

func Algorithms() (algorithms []Algorithm) {
	for _, aa := range asymmetric.Algorithms() {
		algorithms = append(algorithms, aa)
	}
	for _, sa := range symmetric.Algorithms() {
		algorithms = append(algorithms, sa)
	}
	return algorithms
}
