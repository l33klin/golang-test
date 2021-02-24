package benchmark

import (
	"strconv"
	"strings"
	"testing"
)

func stringPlus(a, b string) string {
	var s string
	for i := 1; i < 1000; i++ {
		s += strconv.Itoa(i)
	}
	return s
}

func stringBuild(a, b string) string {
	var builder strings.Builder
	for i := 1; i < 1000; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	return builder.String()
}

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringPlus("hello", "world")
	}
}

func BenchmarkBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuild("hello", "world")
	}
}
