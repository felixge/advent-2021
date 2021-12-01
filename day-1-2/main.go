package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	answer, err := Answer(strings.TrimSpace(string(input)) + "\n")
	fmt.Printf("answer: %v\n", answer)
	return err
}

func Answer(input string) (int, error) {
	var increases int
	var prev []int
	var intVal int
	for _, c := range input {
		if c >= '0' && c <= '9' {
			intVal = intVal*10 + int(c-'0')
		} else if c == '\n' {
			prev = append(prev, intVal)
			if len(prev) >= 4 {
				window := prev[len(prev)-3:]
				prevWindow := prev[len(prev)-4 : len(prev)-1]
				if sum(window) > sum(prevWindow) {
					increases++
				}
			}
			intVal = 0
		} else {
			return 0, fmt.Errorf("bad character in input stream: %q", c)
		}
	}
	return increases, nil
}

func sum(nums []int) (s int) {
	for _, v := range nums {
		s += v
	}
	return
}
