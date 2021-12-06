//go:build v6

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
	return input.Grid().CountEqGt2(), nil
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

func (i *Input) Grid() *Grid {
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

func NewGrid(width, height int) *Grid {
	return &Grid{
		data:   make([]uint8, (width*height+(4-1))/4), // 2 bit per item (ceil int division)
		width:  width,
		height: height,
	}
}

type Grid struct {
	data       []uint8
	countEqGt2 int
	width      int
	height     int
}

func (g *Grid) Get(x, y int) int {
	if x > g.width || y > g.height {
		panic("out of bounds")
	}
	return int(g.get(x, y))
}

func (g *Grid) get(x, y int) uint8 {
	offset, shift := g.offset(x, y)
	return (g.data[offset] >> shift) & 3
}

func (g *Grid) Set(x, y, val int) {
	g.set(x, y, uint8(val))
}

func (g *Grid) set(x, y int, val uint8) {
	offset, shift := g.offset(x, y)
	v := g.data[offset]
	v &^= (3 << shift)
	v |= (val << shift)
	g.data[offset] = v
}

func (g *Grid) Inc(x, y int) {
	offset, shift := g.offset(x, y)
	slotVal := g.data[offset]
	newVal := ((slotVal >> shift) & 3) + 1
	if newVal > 3 {
		return
	}
	slotVal &^= (3 << shift)
	slotVal |= (newVal << shift)
	g.data[offset] = slotVal
	if newVal == 2 {
		g.countEqGt2++
	}
}

func (g *Grid) CountEqGt2() int {
	return g.countEqGt2
}

// 3x2 grid example
//
// x:     0  1  2  0    1  2  -  -
// y:     0  0  0  1    1  2  -  -
// index: 0  0  0  0    1  1  -  -
// shift: 6  4  2  0    6  4  -  -
// data: [00 00 00 00] [00 00 00 00]

func (g *Grid) offset(x, y int) (int, uint8) {
	i := (y*g.width + x)
	return i / 4, 6 - uint8(i)%4*2
}

func (g *Grid) String() string {
	var lines []string
	for y := 0; y < g.height; y++ {
		var line string
		for x := 0; x < g.width; x++ {
			val := g.Get(x, y)
			var char string
			if val == 0 {
				char = "."
			} else {
				char = strconv.Itoa(int(val))
			}
			line += char
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
