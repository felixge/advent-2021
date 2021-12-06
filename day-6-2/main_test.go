package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleInput = strings.TrimSpace(`
3,4,3,1,2
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
		answer, err := Answer(exampleInput)
		require.NoError(t, err)
		require.Equal(t, 26984457539, answer)
	})

	t.Run("puzzle", func(t *testing.T) {
		answer, err := Answer(puzzleInput)
		require.NoError(t, err)
		require.Equal(t, answer, 1681503251694)
	})
}

func BenchmarkAnswer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Answer(puzzleInput)
	}
}
