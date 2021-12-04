//go:build v1

package main

func Answer(input string) (int, error) {
	var numbers []string
	var number string
	for len(input) > 0 {
		if input[0] == '\n' {
			numbers = append(numbers, number)
			number = ""
		} else {
			number += string(input[0])
		}
		input = input[1:]
	}

	oxygen := find(numbers, mostCommon)
	co2 := find(numbers, leastCommon)

	return oxygen * co2, nil
}

func find(numbers []string, filter func([]string, int) byte) int {
	for pos := 0; len(numbers) > 1; pos++ {
		needle := filter(numbers, pos)
		var remaining []string
		for _, v := range numbers {
			if v[pos] == needle {
				remaining = append(remaining, v)
			}
		}
		numbers = remaining
	}
	number := numbers[0]
	decimal := 0
	for i := 0; i < len(number); i++ {
		if number[i] == '1' {
			decimal |= 1 << (len(number) - i - 1)
		}
	}

	return decimal
}

func mostCommon(numbers []string, pos int) byte {
	oneCount := countOnes(numbers, pos)
	zeroCount := len(numbers) - oneCount
	if oneCount >= zeroCount {
		return '1'
	}
	return '0'
}

func leastCommon(numbers []string, pos int) byte {
	oneCount := countOnes(numbers, pos)
	zeroCount := len(numbers) - oneCount
	if oneCount < zeroCount {
		return '1'
	}
	return '0'
}

func countOnes(numbers []string, pos int) int {
	var oneCount int
	for _, v := range numbers {
		if v[pos] == '1' {
			oneCount++
		}
	}
	return oneCount
}
