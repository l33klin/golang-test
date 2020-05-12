package benchmark

import (
	"math/rand"
	"testing"
)

var one int32 = 20

func inc() int32 {
	a := rand.Int31()
	// b := a++
	a++
	return a
}

func addVar() int32 {
	a := rand.Int31()
	a += one
	return a
}

func BenchmarkInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// bit_move()
		inc()
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addVar()
		// mod_100()
	}
}
