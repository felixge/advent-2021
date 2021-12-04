//go:build v1

package main

import (
	"fmt"
	"io"
)

func Answer(input string) (int, error) {
	var g Game
	if err := g.Parse(input); err != nil {
		return 0, err
	}
	return g.Play()
}

type Game struct {
	numbers []int
	boards  []Board
}

func (g *Game) Parse(input string) error {
	const (
		inputDigit = 1
		inputNl    = 2
		boardDigit = 3
		boardSpace = 4
		boardNl    = 5
	)

	var (
		state  = inputDigit
		number = 0
		board  = NewBoard()
	)

	for i := 0; i < len(input); i++ {
		c := input[i]
		switch state {
		case inputDigit:
			if c == ',' || c == '\n' {
				g.numbers = append(g.numbers, number)
				number = 0
				if c == '\n' {
					state = inputNl
				}
			} else if c >= '0' && c <= '9' {
				number = number*10 + int(c) - '0'
			} else {
				return fmt.Errorf("bad inputDigit: %q", string(c))
			}
		case inputNl:
			if c == '\n' {
				state = boardSpace
			} else {
				return fmt.Errorf("bad inputNl: %q", string(c))
			}
		case boardDigit:
			if c == ' ' {
				board.Add(number)
				number = 0
				state = boardSpace
			} else if c == '\n' {
				board.Add(number)
				number = 0
				state = boardNl
			} else if c >= '0' && c <= '9' {
				number = number*10 + int(c) - '0'
				state = boardDigit
			} else {
				return fmt.Errorf("bad boardDigit: %q", string(c))
			}
		case boardSpace:
			if c == ' ' {
				// do nothing
			} else if c >= '0' && c <= '9' {
				number = number*10 + int(c) - '0'
				state = boardDigit
			} else {
				return fmt.Errorf("bad boardNl: %q", string(c))
			}
		case boardNl:
			if c == ' ' {
				state = boardSpace
			} else if c == '\n' {
				g.boards = append(g.boards, board)
				board = NewBoard()
				state = boardSpace
			} else if c >= '0' && c <= '9' {
				number = number*10 + int(c) - '0'
				state = boardDigit
			} else {
				return fmt.Errorf("bad boardNl: %q", string(c))
			}
		}
	}
	g.boards = append(g.boards, board)
	return nil
}

func (g *Game) Play() (int, error) {
	for _, number := range g.numbers {
		for i := range g.boards {
			if result := g.boards[i].Mark(number); result != -1 {
				return result, nil
			}
		}
	}
	return 0, io.EOF
}

const boardSize = 5

func NewBoard() Board {
	return Board{
		numbers: make([]int, 0, boardSize*boardSize),
		marked:  make([]bool, boardSize*boardSize),
	}
}

type Board struct {
	numbers []int
	marked  []bool
}

func (b *Board) Add(number int) {
	b.numbers = append(b.numbers, number)
}

func (b *Board) Mark(number int) int {
	for i, boardNumber := range b.numbers {
		if boardNumber == number {
			b.marked[i] = true
			if b.won(i) {
				return number * b.sumUnmarked()
			}
		}
	}
	return -1
}

func (b *Board) won(i int) bool {
	row, col := b.rowCol(i)
	var cols int
	for checkCol := 0; checkCol < boardSize; checkCol++ {
		if b.marked[b.offset(row, checkCol)] {
			cols++
		}
	}
	if cols == boardSize {
		return true
	}
	var rows int
	for checkRow := 0; checkRow < boardSize; checkRow++ {
		if b.marked[b.offset(checkRow, col)] {
			rows++
		}
	}
	return rows == boardSize
}

func (b *Board) sumUnmarked() (sum int) {
	for i, number := range b.numbers {
		if !b.marked[i] {
			sum += number
		}
	}
	return
}

func (b *Board) rowCol(offset int) (int, int) {
	row := offset / boardSize
	col := offset - row*boardSize
	return row, col
}

func (b *Board) offset(row, col int) int {
	return row*boardSize + col
}
