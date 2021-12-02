package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

var testInput = strings.TrimSpace(`
forward 5
down 5
forward 8
up 3
down 8
forward 2
`)

func TestAnswer(t *testing.T) {
	is := is.New(t)
	answer, err := Answer(testInput)
	is.NoErr(err)
	is.Equal(answer, 900)
}

func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer(testInput)
	}
}
