package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solve(inputs []int) int {
	var total int = 0
	for _, input := range inputs {
		total += (input / 3) - 2
	}
	return total
}

func solvePart2(inputs []int) int {
	var total int = 0
	for _, partMass := range inputs {
		for f := (partMass / 3) - 2; f > 0; f = (f / 3) - 2 {
			total += f
		}
	}
	return total
}

func parseInput(fileContents string) []int {
	var ret []int
	for _, line := range strings.Split(fileContents, "\n") {
		if input, err := strconv.Atoi(line); err != nil {
			panic(err)
		} else {
			ret = append(ret, input)
		}
	}
	return ret
}

func main() {
	fileContents, err := ioutil.ReadFile("src/github.com/superdodd/aoc2019/day1/day1_input.txt")
	if err != nil {
		panic(err)
	}
	inputs := parseInput(string(fileContents))
	solution := solve(inputs)
	solution2 := solvePart2(inputs)
	fmt.Println("Part 1: ", solution)
	fmt.Println("Part 2: ", solution2)
}
