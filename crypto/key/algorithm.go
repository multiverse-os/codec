package key

type AlgorithmType int

const (
	Argon2 AlgorithmType = iota
	Bcrypt
)

type Algorithm struct {
	Type       AlgorithmType
	Parameters map[string]int
}
