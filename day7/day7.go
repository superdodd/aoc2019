package main

import (
	"fmt"
	"github.com/superdodd/aoc2019/common/intcode"
	"sync"
)

var day7_input = []int64{
	3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 55, 80, 97, 118, 199, 280, 361, 442, 99999, 3, 9, 101, 2, 9, 9, 1002, 9, 5, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9,
	101, 5, 9, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 102, 5, 9, 9, 101, 4, 9, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9,
	1002, 9, 2, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 101, 5, 9, 9, 1002, 9, 2, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4,
	9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3,
	9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9,
	1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9,
	102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101,
	1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001,
	9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9,
	4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4,
	9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99,
}

func amplifierSequence(program *intcode.Intcode, phaseSequence []int64) int64 {
	amps := make([]*intcode.Intcode, len(phaseSequence))
	var amp *intcode.Intcode
	for i, p := range phaseSequence {
		//noinspection GoNilness
		amp = amp.Chain(program.Program...)
		amp.Inputs = []int64{p}
		amps[i] = amp
	}
	amps[len(amps)-1].OutputChan = amps[0].InputChan
	// Run the programs in parallel and wait for them to finish.
	var wg sync.WaitGroup
	for _, a := range amps {
		wg.Add(1)
		go func(a *intcode.Intcode) {
			a.MustRun()
			wg.Done()
		}(a)
	}
	// Seed initial amplifier with input of zero
	amps[0].InputChan <- 0
	wg.Wait()
	// The return value is the last output emitted from the last amp.
	lastAmpOutputs := amps[len(amps)-1].Outputs
	return lastAmpOutputs[len(lastAmpOutputs)-1]
}

// Perm calls f with each permutation of a.
func Perm(a []int64, f func([]int64)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int64, f func([]int64), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func solve(program *intcode.Intcode, phases []int64) int64 {
	maxOutput := int64(0)
	Perm(phases, func(phaseSequence []int64) {
		res := amplifierSequence(program, phaseSequence)
		if res > maxOutput {
			maxOutput = res
		}
	})
	return maxOutput
}

func solvePart1(program *intcode.Intcode) int64 {
	return solve(program, []int64{0, 1, 2, 3, 4})
}

func solvePart2(program *intcode.Intcode) int64 {
	return solve(program, []int64{5, 6, 7, 8, 9})
}

func main() {
	program := intcode.NewIntcode(day7_input...)
	fmt.Println("Day 7:")
	fmt.Println("Part 1: ", solvePart1(program))
	fmt.Println("Part 2: ", solvePart2(program))
}
