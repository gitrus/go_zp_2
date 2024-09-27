package randgen_test

import (
	"go_zp_2/pkg/randgen"
	"testing"
)

func BenchmarkCryptoRand128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := randgen.CryptoRand128()
		if err != nil {
			b.Fatalf("crypto/rand error: %v", err)
		}
	}
}

func BenchmarkMathRand128(b *testing.B) {
	biGen := randgen.NewBigIntRandom(123)
	for i := 0; i < b.N; i++ {
		_, _ = biGen.MathRand128()
	}
}
