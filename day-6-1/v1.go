//go:build v1

package main

import (
	"strconv"
	"strings"
)

func Answer(inputText string) (int, error) {
	var school []int
	for _, ageS := range strings.Split(strings.TrimSpace(inputText), ",") {
		age, err := strconv.Atoi(ageS)
		if err != nil {
			return 0, err
		}
		school = append(school, age)
	}

	for i := 0; i < 80; i++ {
		for i, age := range school {
			school[i]--
			if age == 0 {
				school[i] = 6
				school = append(school, 8)
			}
		}
	}
	return len(school), nil
}
