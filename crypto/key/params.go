package key

////////////////////////////////////////////////////////////////////////////////

// PARAMS //////////////////////////////////////////////////////////////////////
//func (self CryptoSystem) String(name string) string { return self.Params.String(name) }
//func (self CryptoSystem) Integer(name string) int   { return self.Params.Integer(name) }
//func (self CryptoSystem) Float(name string) float64 { return self.Params.Float(name) }
//func (self CryptoSystem) Boolean(name string) bool  { return self.Params.Boolean(name) }
//
//func (self CryptoSystem) AddString(name string, value string) CryptoSystem {
//	self.Params.AddString(name, value)
//	return self
//}
//
//func (self CryptoSystem) AddInteger(name string, value int) CryptoSystem {
//	self.Params.AddInteger(name, value)
//	return self
//}
//
//func (self CryptoSystem) AddFloat(name string, value float64) CryptoSystem {
//	self.Params.AddFloat(name, value)
//	return self
//}
//
//func (self CryptoSystem) AddBoolean(name string, value bool) CryptoSystem {
//	self.Params.AddBoolean(name, value)
//	return self
//}

////////////////////////////////////////////////////////////////////////////////

// type SymmetricKey interface {
// 	Key
//
// 	SecretKey() []byte
// }
//
// type AsymmetricKey interface {
// 	Key
// 	Sign(input []byte) ([]byte, error)
// 	Verify(input []byte) (bool, error)
//
// 	PublicKey() []byte
// 	PrivateKey() []byte
// }
//
// ////////////////////////////////////////////////////////////////////////////////
// type Key interface {
// 	String() string
//
// 	Encrypt(plainText []byte) ([]byte, error)
// 	Decrypt(cipherText []byte) ([]byte, error)
// }
