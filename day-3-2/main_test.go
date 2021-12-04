package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/matryer/is"
)

var exampleInput = strings.TrimSpace(`
00100s
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
`) + "\n"

var puzzleInput = (func() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data)) + "\n"
})()

func TestAnswer(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		is := is.New(t)
		answer, err := Answer(exampleInput)
		is.NoErr(err)
		is.Equal(answer, 230)
	})

	t.Run("puzzle", func(t *testing.T) {
		is := is.New(t)
		answer, err := Answer(puzzleInput)
		is.NoErr(err)
		is.Equal(answer, 2784375)
	})
}

func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer(exampleInput)
	}
}
