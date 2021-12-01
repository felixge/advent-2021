package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

var testInput = strings.TrimSpace(`
199
200
208
210
200
207
240
269
260
263
`) + "\n"

func TestAnswer(t *testing.T) {
	is := is.New(t)
	answer, err := Answer(testInput)
	is.NoErr(err)
	is.Equal(answer, 5)
}

func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer(testInput)
	}
}
