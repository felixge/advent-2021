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
	answer, err := Answer(strings.TrimSpace(string(input)))
	fmt.Printf("answer: %v\n", answer)
	return err
}

func Answer(input string) (int, error) {
	var prev struct {
		val int64
		set bool
	}
	var increases int
	var line string
	for len(input) > 0 {
		i := strings.IndexByte(input, '\n')
		if i == -1 {
			line = input
			input = ""
		} else {
			line = input[0:i]
			input = input[i+1:]
		}
		val, err := parseInt(line)
		if err != nil {
			return 0, err
		}
		if val > prev.val && prev.set {
			increases++
		}
		prev.set = true
		prev.val = val
	}
	return increases, nil
}

func parseInt(val string) (int64, error) {
	var intVal int64
	factor := int64(1)
	for i := len(val) - 1; i >= 0; i-- {
		c := val[i]
		if c >= '0' && c <= '9' {
			intVal += int64(c-'0') * factor
		} else {
			return intVal, fmt.Errorf("bad int: %q", val)
		}
		factor *= 10
	}
	return intVal, nil
}
