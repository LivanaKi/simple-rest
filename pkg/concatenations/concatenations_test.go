package concatenations_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Users/natza/simple-rest/pkg/concatenations"
)

func GenerateString() []string {
	n := rand.Intn(71) + 30 //nolint

	bs := make([]string, n)
	for i := 0; i < n; i++ {
		num := 65 + rand.Intn(26) //nolint
		bs[i] = strconv.Itoa(num)
	}
	return bs
}

var str = GenerateString()

func BenchmarkConcatOne(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		concatenations.ConcatOne(str)
	}
}

func BenchmarkConcatTwo(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		concatenations.ConcatTwo(str)
	}
}

func BenchmarkConcatThree(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		concatenations.ConcatThree(str)
	}
}
