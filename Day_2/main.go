package main

import (
	"fmt"
	"strconv"
)

func main() {
	idRanges, err := parseFile("./input.txt")
	if err != nil {
		return
	}

	invalidIdSum := getInvalidIdSum(idRanges)

	fmt.Println("The sum of the invalid ids is: ", invalidIdSum)
}

func getInvalidIdSum(idRanges []IDRange) int {
	sum := 0

	for _, idRange := range idRanges {
		sum += getInvalidIdSumForRange(idRange.Lower, idRange.Upper)
	}

	return sum
}

func getInvalidIdSumForRange(lowerRange int, upperRange int) int {
	sum := 0

	for i := lowerRange; i <= upperRange; i++ {
		num := strconv.Itoa(i)

		if isInvalidIdPattern(num) {
			sum += i
		}
	}

	return sum
}

func isInvalidIdPattern(numStr string) bool {
	if len(numStr) <= 1 {
		return false
	}

	for i := 0; i < len(numStr); i++ {
		pattern := numStr[:i]

		if len(pattern) > 0 && isRecurringPattern(numStr, pattern) {
			return true
		}
	}

	return false
}

func isRecurringPattern(checkStr string, pattern string) bool {
	for i := 0; i < len(checkStr); i++ {
		for j := 0; j < len(pattern); j++ {
			if i < len(checkStr) && checkStr[i] == checkStr[j] {
				i++
			} else {
				return false
			}
		}
		i--
	}
	return true
}
