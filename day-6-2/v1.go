//go:build v1

package main

import (
	"strconv"
	"strings"
)

func Answer(inputText string) (int, error) {
	school := make([]int, 9)
	for _, ageS := range strings.Split(strings.TrimSpace(inputText), ",") {
		age, err := strconv.Atoi(ageS)
		if err != nil {
			return 0, err
		}
		school[age]++
	}

	for i := 0; i < 256; i++ {
		newFish := school[0]
		copy(school, school[1:])
		school[6] += newFish
		school[8] = newFish
	}
	var total int
	for _, count := range school {
		total += count
	}
	return total, nil
}
