package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/matryer/is"
)

var exampleInput = strings.TrimSpace(`
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`) + "\n"

var puzzleInput = (func() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data)) + "\n"
})()

func TestGame(t *testing.T) {
	t.Run("parse", func(t *testing.T) {
		is := is.New(t)
		var g Game
		is.NoErr(g.Parse(exampleInput))

		is.Equal(len(g.numbers), 27)
		is.Equal(g.numbers[0], 7)
		is.Equal(g.numbers[len(g.numbers)-1], 1)

		is.Equal(len(g.boards), 3)
		for _, b := range g.boards {
			is.Equal(len(b.numbers), 25)
		}
		is.Equal(g.boards[0].numbers[0], 22)
		is.Equal(g.boards[0].numbers[len(g.boards[0].numbers)-1], 19)
		is.Equal(g.boards[1].numbers[0], 3)
		is.Equal(g.boards[1].numbers[len(g.boards[0].numbers)-1], 6)
		is.Equal(g.boards[2].numbers[0], 14)
		is.Equal(g.boards[2].numbers[len(g.boards[0].numbers)-1], 7)
	})
}

func TestAnswer(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		is := is.New(t)
		answer, err := Answer(exampleInput)
		is.NoErr(err)
		is.Equal(answer, 1924)
	})

	t.Run("puzzle", func(t *testing.T) {
		is := is.New(t)
		answer, err := Answer(puzzleInput)
		is.NoErr(err)
		is.Equal(answer, 12738)
	})
}

func BenchmarkAnswer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Answer(puzzleInput)
	}
}
