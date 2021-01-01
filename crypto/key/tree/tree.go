package key

//type Identity struct {
//	Tree *Tree
//
//	Subidentities []*Identity
//}

// TODO: Or should it be key.Graph? Probably graph since we want to model
// myclellium. Growing out in all directions
type AccessType int

const (
	RootKey AccessType = iota
	SessionKey
)

type Tree struct {
	Root *Keypair

	Keys         map[int64]*KeyPair
	Tokens       map[int64]*Token
	Certificates map[int64]*Certificate
}

//type RecoveryType int
//
//// NOTE: Need two (2) of three (3) of the recovery methods to reach the required
//// threshold to rebuild your key tree, and reset your passowrd while maintaining
//// your identity.
//const (
//	OfflineSecurityQuestion RecoveryType = iota
//	RootEncryptedBackup
//	// NOTE: Split backup can be stored offline, on paper, or on friends/family
//	// machines, or on smart card.
//	SplitBackup
//)

//type Point struct {
//	X int
//	Y int
//}
