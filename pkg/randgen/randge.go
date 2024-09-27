package randgen

import (
	cryptoRand "crypto/rand"
	"math"
	"math/big"
	mathRand "math/rand"
)

func CryptoRand128() (*big.Int, error) {
	// 128 -> 2^7 -> 2^3 * 2^4 -> 16 bytes UUID == int128
	buf := make([]byte, 16)

	_, err := cryptoRand.Read(buf)
	if err != nil {
		return nil, err
	}

	result := new(big.Int).SetBytes(buf)

	return result, nil
}

type BigIntRandom struct {
	randSrc mathRand.Source
}

func NewBigIntRandom(seed int64) *BigIntRandom {
	return &BigIntRandom{
		randSrc: mathRand.NewSource(seed),
	}
}
func (bir *BigIntRandom) MathRand128() (*big.Int, error) {
	// 128 -> 2^7 -> 2^3 * 2^4 -> 16 bytes
	buf := make([]byte, 16)

	r := mathRand.New(bir.randSrc)

	// Fill the buffer with random bytes
	for i := range buf {
		buf[i] = byte(r.Intn(math.MaxUint8 + 1)) // Random byte values in range [0, 255]
	}

	// Convert the buffer into a big.Int
	result := new(big.Int).SetBytes(buf)

	return result, nil
}
