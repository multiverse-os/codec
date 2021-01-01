package key

type (
	Public  []byte
	Private []byte
	Hash    []byte
)

type Key interface {
	PublicKey() []byte
	PrivateKey() []byte

	//Deterministic(algorithm cryptography.Algorithm, index []Point) *Key

	// Tree
	Root() *Key
	Parent() *Key
	Subkeys() []*Key

	GenerateSubkey() *Key
}
