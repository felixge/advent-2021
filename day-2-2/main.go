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
	var aim, x, y int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return 0, fmt.Errorf("bad line: %q", line)
		}
		amount, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("bad amount: %q: %s", parts, err)
		}
		switch parts[0] {
		case "forward":
			x += int(amount)
			y += aim * int(amount)
		case "down":
			aim += int(amount)
		case "up":
			aim -= int(amount)
		default:
			return 0, fmt.Errorf("bad command: %q", line)
		}
	}
	return x * y, nil
}
