package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unsafe"
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

// Solution from Valentin Deleplace: https://twitter.com/val_deleplace/status/1466442802330488838

func Answer(input string) (int, error) {
	var prev int64 = -9999
	var increases int = -1

	val := int64(0)
	for p, N := 0, len(input); p < N; p++ {
		c := input[p]
		notNl := boolToInt64(c != '\n')
		increases += int((^notNl & 1) * boolToInt64(val > prev))
		prev = notNl*prev + val*(^notNl&1)
		val = notNl * ((val << 8) + int64(c))
	}
	increases += int(boolToInt64(val > prev))
	return increases, nil
}

func boolToInt64(b bool) int64 {
	return *(*int64)(unsafe.Pointer(&b)) & 1
}
