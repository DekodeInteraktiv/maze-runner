package game

type ClaimType uint16

const (
	Unclaimed ClaimType = iota
	Red
	Blue
	Green
	Yellow
)

func newClaims(size int) [][]ClaimType {
	claims := make([][]ClaimType, size)
	for i := range claims {
		claims[i] = make([]ClaimType, size)
	}

	return claims
}
