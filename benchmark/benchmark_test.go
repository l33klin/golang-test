package benchmark

import (
	"math/rand"
	"testing"
)

func mod_100() int32 {
	a := rand.Int31()
	b := a % 100
	return b
}

func bit_move() int32 {
	a := rand.Int31()
	b := a << 26
	return b
}

func BenchmarkMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// bit_move()
		mod_100()
	}
}

func BenchmarkMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bit_move()
		// mod_100()
	}
}
