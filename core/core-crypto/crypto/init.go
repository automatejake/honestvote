package crypto

import (
	"crypto/elliptic"
	"math/big"
)

var zero *big.Int
var two *big.Int
var p256 elliptic.Curve

func init() {
	zero = big.NewInt(0)
	two = big.NewInt(2)
	p256 = elliptic.P256()
}
