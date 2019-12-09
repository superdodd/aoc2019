package main

import (
	"fmt"
	"github.com/superdodd/aoc2019/common/intcode"
)

var day2_program = []int64{
	1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 2, 23, 6, 27, 2, 6, 27, 31, 2, 13, 31,
	35, 1, 10, 35, 39, 2, 39, 13, 43, 1, 43, 13, 47, 1, 6, 47, 51, 1, 10, 51, 55, 2, 55, 6, 59, 1, 5, 59, 63, 2, 9, 63,
	67, 1, 6, 67, 71, 2, 9, 71, 75, 1, 6, 75, 79, 2, 79, 13, 83, 1, 83, 10, 87, 1, 13, 87, 91, 1, 91, 10, 95, 2, 9, 95,
	99, 1, 5, 99, 103, 2, 10, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0,
}

func solvePart1(ic *intcode.Intcode) int64 {
	// First, modify the input program
	ic.Reset()
	ic.Program[1] = 12
	ic.Program[2] = 2
	ic.MustRun()
	return ic.Program[0]
}

func solvePart2(ic *intcode.Intcode) int64 {
	for noun := int64(0); noun <= 99; noun++ {
		for verb := int64(0); verb <= 99; verb++ {
			ic.Reset()
			ic.Program[1] = noun
			ic.Program[2] = verb
			err := ic.Run()
			if ic.Program[0] == 19690720 && err == nil {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func main() {
	ic := intcode.NewIntcode(day2_program...)

	/*
		Day 2
		Part 1:  5434663
		Part 2:  4559
	*/
	fmt.Println("Day 2")
	fmt.Println("Part 1: ", solvePart1(ic))
	fmt.Println("Part 2: ", solvePart2(ic))
}
