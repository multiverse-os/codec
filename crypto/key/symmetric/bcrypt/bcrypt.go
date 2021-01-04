package bcrypt

import (
	bcrypt "golang.org/x/crypto/bcrypt"
)

type Crypto struct{}

func CryptoHandler() *Crypto {
	return &Crypto{}
}

type CryptoInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

func (crp *Crypto) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (crp *Crypto) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
