package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	nBig, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, two))
	return nBig.Add(nBig, two)
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
