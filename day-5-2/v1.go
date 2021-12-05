//go:build v1

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Answer(inputText string) (int, error) {
	input, err := ParseInput(inputText)
	if err != nil {
		return 0, err
	}
	return input.Grid().Count(func(val int) bool {
		return val >= 2
	}), nil
}

func ParseInput(text string) (*Input, error) {
	// parser states
	const (
		x1 = iota
		y1
		arrowSpace
		arrowDash
		arrowGt
		x2
		y2
	)

	var (
		input Input
		state = x1
		line  Line
	)

	for i := 0; i < len(text); i++ {
		c := text[i]
		switch state {
		case x1:
			if c >= '0' && c <= '9' {
				line.X1 = line.X1*10 + int(c) - '0'
			} else if c == ',' {
				state = y1
			} else {
				return nil, fmt.Errorf("x1: bad char: %q", string(c))
			}
		case y1:
			if c >= '0' && c <= '9' {
				line.Y1 = line.Y1*10 + int(c) - '0'
			} else if c == ' ' {
				state = arrowSpace
			} else {
				return nil, fmt.Errorf("y1: bad char: %q", string(c))
			}
		case arrowSpace:
			if c == '-' {
				state = arrowDash
			} else {
				return nil, fmt.Errorf("arrowSpace: bad char: %q", string(c))
			}
		case arrowDash:
			if c == '>' {
				state = arrowGt
			} else {
				return nil, fmt.Errorf("arrowDash: bad char: %q", string(c))
			}
		case arrowGt:
			if c == ' ' {
				state = x2
			} else {
				return nil, fmt.Errorf("arrowGt: bad char: %q", string(c))
			}
		case x2:
			if c >= '0' && c <= '9' {
				line.X2 = line.X2*10 + int(c) - '0'
			} else if c == ',' {
				state = y2
			} else {
				return nil, fmt.Errorf("x2: bad char: %q", string(c))
			}
		case y2:
			if c >= '0' && c <= '9' {
				line.Y2 = line.Y2*10 + int(c) - '0'
			} else if c == '\n' {
				input.lines = append(input.lines, line)
				line = Line{}
				state = x1
			} else {
				return nil, fmt.Errorf("y2: bad char: %q", string(c))
			}
		}
	}
	return &input, nil
}

type Input struct {
	lines []Line
}

func (i *Input) Grid() Grid {
	var (
		width  = -1
		height = -1
	)
	for _, l := range i.lines {
		if l.X1 > width {
			width = l.X1
		}
		if l.X2 > width {
			width = l.X2
		}
		if l.Y1 > height {
			height = l.Y1
		}
		if l.Y2 > height {
			height = l.Y2
		}
	}

	g := NewGrid(width+1, height+1)
	for _, l := range i.lines {
		x, y := l.X1, l.Y1
		for {
			g.Inc(x, y)
			if l.X1 < l.X2 {
				x++
				if x > l.X2 {
					break
				}
			} else if l.X1 > l.X2 {
				x--
				if x < l.X2 {
					break
				}
			}

			if l.Y1 < l.Y2 {
				y++
				if y > l.Y2 {
					break
				}
			} else if l.Y1 > l.Y2 {
				y--
				if y < l.Y2 {
					break
				}
			}
		}
	}

	return g
}

type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func NewGrid(width, height int) Grid {
	return Grid{
		data:   make([]int, width*height),
		width:  width,
		height: height,
	}
}

type Grid struct {
	data   []int
	width  int
	height int
}

func (g Grid) Get(x, y int) int {
	return g.data[g.offset(x, y)]
}

func (g Grid) Set(x, y, val int) {
	g.data[g.offset(x, y)] = val
}

func (g Grid) Inc(x, y int) {
	g.data[g.offset(x, y)]++
}

func (g Grid) Count(fn func(val int) bool) (count int) {
	for _, v := range g.data {
		if fn(v) {
			count++
		}
	}
	return
}

func (g Grid) offset(x, y int) int {
	return y*g.width + x
}

func (g Grid) String() string {
	var lines []string
	for y := 0; y < g.height; y++ {
		var line string
		for x := 0; x < g.width; x++ {
			val := g.Get(x, y)
			var char string
			if val == 0 {
				char = "."
			} else {
				char = strconv.Itoa(val)
			}
			line += char
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
