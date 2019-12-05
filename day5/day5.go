package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func parseInput(fileContents string) []int {
	var ret []int
	for _, i := range strings.Split(fileContents, ",") {
		ival, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ret = append(ret, ival)
	}
	return ret
}

// Return a value representing either the immediate or positional version of a parameter.
func modeValue(input []int, pc int, index int) int {
	if (input[pc]/(int(math.Pow10(index+1))))%10 == 0 {
		return input[input[pc+index]]
	}
	return input[pc+index]
}

func runProgram(progMem []int, userInput int) error {
	pc := 0
	for progMem[pc] != 99 {
		switch progMem[pc] % 100 {
		case 1: // Add
			progMem[progMem[pc+3]] = modeValue(progMem, pc, 1) + modeValue(progMem, pc, 2)
			pc += 4
		case 2: // Multiply
			progMem[progMem[pc+3]] = modeValue(progMem, pc, 1) * modeValue(progMem, pc, 2)
			pc += 4
		case 3: // Input
			progMem[progMem[pc+1]] = userInput
			pc += 2
		case 4: // Output
			out := modeValue(progMem, pc, 1)
			fmt.Println(out)
			pc += 2
		case 5: // Jump if true
			if modeValue(progMem, pc, 1) != 0 {
				pc = modeValue(progMem, pc, 2)
			} else {
				pc += 3
			}
		case 6: // Jump if false
			if modeValue(progMem, pc, 1) == 0 {
				pc = modeValue(progMem, pc, 2)
			} else {
				pc += 3
			}
		case 7: // Less than
			if modeValue(progMem, pc, 1) < modeValue(progMem, pc, 2) {
				progMem[progMem[pc+3]] = 1
			} else {
				progMem[progMem[pc+3]] = 0
			}
			pc += 4
		case 8: // Equals
			if modeValue(progMem, pc, 1) == modeValue(progMem, pc, 2) {
				progMem[progMem[pc+3]] = 1
			} else {
				progMem[progMem[pc+3]] = 0
			}
			pc += 4
		default: // Error
			return fmt.Errorf("unexpected opcode at pc=%d: %d", pc, progMem[pc])
		}
	}
	return nil
}

func main() {
	fileContents, err := ioutil.ReadFile("src/github.com/superdodd/aoc2019/day5/day5_input.txt")
	if err != nil {
		panic(err)
	}
	inputs := parseInput(string(fileContents))
	fmt.Println("Day 5")
	fmt.Println("Part 1:")
	part1Program := append([]int(nil), inputs...)
	if err = runProgram(part1Program, 1); err != nil {
		panic(err)
	}
	part2Program := append([]int(nil), inputs...)
	fmt.Println("Part 2:")
	if err = runProgram(part2Program, 5); err != nil {
		panic(err)
	}
}
