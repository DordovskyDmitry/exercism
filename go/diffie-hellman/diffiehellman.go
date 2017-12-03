package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var one = big.NewInt(1)
var two = big.NewInt(2)

func PrivateKey(p *big.Int) *big.Int {
	if p.Cmp(one) == 0 {
		return p
	}

	nBig, err := rand.Int(rand.Reader, new(big.Int).Sub(p, two))

	if err != nil {
		panic(err)
	}

	if nBig.Cmp(two) < 0 {
		nBig.Add(nBig, two)
	}

	return nBig
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	x := big.NewInt(g)
	return new(big.Int).Exp(x, private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return

}
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
