package main

import (
	"fmt"
	"math"
)

var day12_input = []int{1, -4, 3, -14, 9, -4, -4, -6, 7, 6, -9, -11}

func updateVel(pos, vel []int) {
	for i := range pos {
		for m := 0; m < 4; m++ {
			if pos[3*m+(i%3)] < pos[i] {
				vel[i]--
			}
			if pos[3*m+(i%3)] > pos[i] {
				vel[i]++
			}
		}
	}
}

func updatePos(pos, vel []int) {
	for i := range pos {
		pos[i] += vel[i]
	}
}

func energy(values []int) []int {
	var ret = make([]int, 4)
	for m := 0; m < 4; m++ {
		for i := 0; i < 3; i++ {
			ret[m] += int(math.Abs(float64(values[3*m+i])))
		}
	}
	return ret
}

func totalEnergy(pos, vel []int) int {
	pot := energy(pos)
	kin := energy(vel)
	var ret int
	for i := range pot {
		ret += pot[i] * kin[i]
	}
	return ret
}

func step(pos, vel []int) int {
	updateVel(pos, vel)
	updatePos(pos, vel)
	return totalEnergy(pos, vel)
}

func main() {
	fmt.Println("Day 12:")
	totalEnergy := solvePart1(append([]int(nil), day12_input...))
	fmt.Println("Part 1: ", totalEnergy)
	fmt.Println("Part 2: ", solvePart2(append([]int(nil), day12_input...)))
}

func solvePart1(pos []int) int {
	vel := make([]int, len(pos))

	var totalEnergy int
	for i := 0; i < 1000; i++ {
		totalEnergy = step(pos, vel)
	}
	return totalEnergy
}

func sliceEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func solvePart2BruteForce(pos []int) int {
	initial := append([]int(nil), pos...)
	vel := make([]int, len(pos))
	step(pos, vel)
	var i int
	for i = 1; !sliceEqual(pos, initial); i++ {
		step(pos, vel)
	}
	return i
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solvePart2(pos []int) int {
	opos := append([]int(nil), pos...)
	vel := make([]int, len(pos))
	var x, y, z int
	for s := 1; x == 0 || y == 0 || z == 0; s++ {
		step(pos, vel)
		if x == 0 && pos[0] == opos[0] && pos[3] == opos[3] && pos[6] == opos[6] && pos[9] == opos[9] && vel[0] == 0 && vel[3] == 0 && vel[6] == 0 && vel[9] == 0 {
			x = s
		}
		if y == 0 && pos[1] == opos[1] && pos[4] == opos[4] && pos[7] == opos[7] && pos[10] == opos[10] && vel[1] == 0 && vel[4] == 0 && vel[7] == 0 && vel[10] == 0 {
			y = s
		}
		if z == 0 && pos[2] == opos[2] && pos[5] == opos[5] && pos[8] == opos[8] && pos[11] == opos[11] && vel[2] == 0 && vel[5] == 0 && vel[8] == 0 && vel[11] == 0 {
			z = s
		}
	}
	lcm := x * y / gcd(x, y)
	return lcm * z / gcd(lcm, z)
}
