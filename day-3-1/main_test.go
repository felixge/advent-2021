package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

var testInput = strings.TrimSpace(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)

func TestAnswer(t *testing.T) {
	is := is.New(t)
	answer, err := Answer(testInput)
	is.NoErr(err)
	is.Equal(answer, 198)
}

func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer(testInput)
	}
}
