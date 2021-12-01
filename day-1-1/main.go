package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	for _, line := range strings.Split(input, "\n") {
		val, err := strconv.ParseInt(line, 10, 64)
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
