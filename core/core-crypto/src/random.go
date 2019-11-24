package random

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

func init() {
	var b [16]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func NextRandomNum() float64 {
	return math_rand.Float64()
}
