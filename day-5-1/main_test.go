package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleInput = strings.TrimSpace(`
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`) + "\n"

var puzzleInput = (func() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data)) + "\n"
})()

func TestGrid(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		g := NewGrid(3, 4)
		require.Equal(t, 0, g.Get(0, 0))
		require.Equal(t, 0, g.Get(2, 3))
		require.Panics(t, func() {
			g.Get(2, 4)
		})
		require.Panics(t, func() {
			g.Get(3, 3)
		})
	})

	t.Run("Set", func(t *testing.T) {
		g := NewGrid(3, 4)
		g.Set(0, 0, 2)
		g.Set(2, 3, 3)
		require.Equal(t, 2, g.Get(0, 0))
		require.Equal(t, 3, g.Get(2, 3))
	})

	t.Run("String", func(t *testing.T) {
		g := NewGrid(2, 3)
		require.Equal(t, "..\n..\n..", g.String())
		g.Set(0, 0, 3)
		g.Set(0, 1, 7)
		g.Set(1, 2, 5)
		require.Equal(t, "3.\n7.\n.5", g.String())
	})
}

func BenchmarkGrid(b *testing.B) {
	g := NewGrid(10, 15)
	for i := 0; i < b.N; i++ {
		g.Set(i%10, i%15, i)
		_ = g.Get(i%10, i%15)
	}
}

func TestInput(t *testing.T) {
	t.Run("ParseInput", func(t *testing.T) {
		input, err := ParseInput(exampleInput)
		require.NoError(t, err)

		require.Equal(t, 0, input.lines[0].X1)
		require.Equal(t, 9, input.lines[0].Y1)
		require.Equal(t, 5, input.lines[0].X2)
		require.Equal(t, 9, input.lines[0].Y2)

		require.Equal(t, 5, input.lines[9].X1)
		require.Equal(t, 5, input.lines[9].Y1)
		require.Equal(t, 8, input.lines[9].X2)
		require.Equal(t, 2, input.lines[9].Y2)

		input, err = ParseInput(puzzleInput)
		require.NoError(t, err)
		require.Equal(t, 957, input.lines[0].X1)
		require.Equal(t, 596, input.lines[0].Y1)
		require.Equal(t, 957, input.lines[0].X2)
		require.Equal(t, 182, input.lines[0].Y2)
	})

	t.Run("Grid", func(t *testing.T) {
		input, err := ParseInput(exampleInput)
		require.NoError(t, err)
		grid := input.Grid()
		want := strings.TrimSpace(`
.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
`)
		require.Equal(t, want, grid.String())
	})
}

func TestAnswer(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		answer, err := Answer(exampleInput)
		require.NoError(t, err)
		require.Equal(t, 5, answer)
	})

	t.Run("puzzle", func(t *testing.T) {
		answer, err := Answer(puzzleInput)
		require.NoError(t, err)
		require.Equal(t, answer, 6311)
	})
}

func BenchmarkAnswer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Answer(puzzleInput)
	}
}
