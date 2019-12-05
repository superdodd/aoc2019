package main

import (
	"fmt"
	"math"
)

func validatePassword(password int) bool {
	if password < 100000 || password > 999999 {
		return false
	}
	digits := make([]int, 6)
	for i := range digits {
		digits[i] = (password / int(math.Pow10(5-i))) % 10
	}
	var foundRepeat bool
	for i := 0; i < len(digits)-1; i++ {
		if digits[i+1] == digits[i] {
			foundRepeat = true
		}
		if digits[i+1] < digits[i] {
			return false
		}
	}
	return foundRepeat
}

func solvePart1() int {
	count := 0
	for p := 278384; p <= 824795; p++ {
		if validatePassword(p) {
			count++
		}
	}
	return count
}

func validatePassword2(password int) bool {
	// Assume the password otherwise meets requirements.  Then ensure that we see a repeat of
	// exactly two digits.
	digits := make([]int, 6)
	for i := range digits {
		digits[i] = (password / int(math.Pow10(5-i))) % 10
	}
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			if i > 0 && digits[i-1] == digits[i] {
				continue
			}
			if i < len(digits)-2 && digits[i+1] == digits[i+2] {
				continue
			}
			return true
		}
	}
	return false
}

func solvePart2() int {
	count := 0
	for p := 278384; p <= 824975; p++ {
		if validatePassword(p) && validatePassword2(p) {
			count++
		}
	}
	return count
}

func main() {
	/*
		Day 4:
		Part 1: 921
		Part 2: 603
	*/
	fmt.Println("Day 4:")
	fmt.Println("Part 1: ", solvePart1())
	fmt.Println("Part 2: ", solvePart2())
}
