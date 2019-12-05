package main

import (
	"fmt"
	"io/ioutil"
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

func runProgram(input []int) error {
	pc := 0
loop:
	for {
		switch input[pc] {
		case 1: // Add
			input[input[pc+3]] = input[input[pc+1]] + input[input[pc+2]]
			pc += 4
		case 2: // Multiply
			input[input[pc+3]] = input[input[pc+1]] * input[input[pc+2]]
			pc += 4
		case 99: // Terminate
			break loop
		default: // Error
			return fmt.Errorf("unexpected opcode at pc=%d: %d", pc, input[pc])
		}
	}
	return nil
}

func solvePart1(input []int) int {
	// First, modify the input
	input[1] = 12
	input[2] = 2
	if err := runProgram(input); err != nil {
		panic(err)
	}
	return input[0]
}

func solvePart2(input []int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			testInput := append([]int(nil), input...)
			testInput[1] = noun
			testInput[2] = verb
			err := runProgram(testInput)
			if testInput[0] == 19690720 && err == nil {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func main() {
	fileContents, err := ioutil.ReadFile("src/github.com/superdodd/aoc2019/day2/day2_input.txt")
	if err != nil {
		panic(err)
	}
	inputs := parseInput(string(fileContents))
	solution := solvePart1(append([]int(nil), inputs...))
	solution2 := solvePart2(append([]int(nil), inputs...))
	fmt.Println("Day 2")
	fmt.Println("Part 1: ", solution)
	fmt.Println("Part 2: ", solution2)
}
