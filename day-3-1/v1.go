//go:build v1

package main

import "fmt"

func Answer(input string) (int, error) {
	var (
		col    = 0
		counts []int
		lines  = 0
	)
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch c {
		case '0':
			if col >= len(counts) {
				counts = append(counts, 0)
			}
			col++
		case '1':
			if col >= len(counts) {
				counts = append(counts, 1)
			} else {
				counts[col]++
			}
			col++
		case '\n':
			col = 0
			lines++
		default:
			return 0, fmt.Errorf("bad char: %q", string(c))
		}
	}

	var gamma, epsilon int
	for i, v := range counts {
		if v > lines/2 {
			gamma = gamma | (1 << (len(counts) - i - 1))
		} else {
			epsilon = epsilon | (1 << (len(counts) - i - 1))
		}
	}
	return gamma * epsilon, nil
}
