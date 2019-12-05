package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func (c Coord) Distance(other Coord) int {
	var ret int
	if c.x < other.x {
		ret += other.x - c.x
	} else {
		ret += c.x - other.x
	}
	if c.y < other.y {
		ret += other.y - c.y
	} else {
		ret += c.y - other.y
	}
	return ret
}

func parse(fileContents []byte) ([][]Coord, error) {
	var ret [][]Coord
	for _, line := range strings.Split(string(fileContents), "\n") {
		var wire []Coord
		pos := Coord{}
		for _, step := range strings.Split(line, ",") {
			dir := step[0]
			count, err := strconv.Atoi(step[1:])
			if err != nil {
				return nil, err
			}
			for i := 0; i < count; i++ {
				switch dir {
				case 'U':
					pos.y += 1
				case 'D':
					pos.y -= 1
				case 'L':
					pos.x -= 1
				case 'R':
					pos.x += 1
				}
				wire = append(wire, pos)
			}
		}
		ret = append(ret, wire)
	}
	return ret, nil
}

func solvePart1(wires [][]Coord) int {
	wiresMap, maxDist := buildWiresMap(wires)

	for d := 0; d < maxDist; d++ {
		for x := 0; x <= d; x++ {
			y := d - x
			if checkIntersection(wiresMap, x, y) || checkIntersection(wiresMap, -x, y) ||
				checkIntersection(wiresMap, x, -y) || checkIntersection(wiresMap, -x, -y) {
				return Coord{0, 0}.Distance(Coord{x, y})
			}
		}
	}
	return -1
}

func solvePart2(wires [][]Coord) int {
	wiresMap, _ := buildWiresMap(wires)

	minDist := 0
	for i, c := range wires[0] {
		if checkIntersection(wiresMap, c.x, c.y) {
			totalSteps := len(wiresMap)
			for j, w := range wires {
				if j == 0 {
					totalSteps += i
				} else {
					for s, c2 := range w {
						if c.x == c2.x && c.y == c2.y {
							totalSteps += s
							break
						}
					}
				}
			}
			if minDist == 0 || totalSteps < minDist {
				minDist = totalSteps
			}
		}
	}
	return minDist
}

func buildWiresMap(wires [][]Coord) ([]map[int]map[int]struct{}, int) {
	var wiresMap []map[int]map[int]struct{}
	var maxDist int
	for _, wire := range wires {
		wireMap := map[int]map[int]struct{}{}
		for _, c := range wire {
			d := c.Distance(Coord{0, 0})
			if d > maxDist {
				maxDist = d
			}
			yMap, present := wireMap[c.x]
			if !present {
				yMap = make(map[int]struct{})
			}
			yMap[c.y] = struct{}{}
			wireMap[c.x] = yMap
		}
		wiresMap = append(wiresMap, wireMap)
	}
	return wiresMap, maxDist
}

func checkIntersection(wiresMap []map[int]map[int]struct{}, x int, y int) bool {
	for _, wireMap := range wiresMap {
		yMap, xPresent := wireMap[x]
		if !xPresent {
			return false
		}
		_, yPresent := yMap[y]
		if !yPresent {
			return false
		}
	}
	return true
}

func main() {
	fileContents, err := ioutil.ReadFile("src/github.com/superdodd/aoc2019/day3/day3_input.txt")
	if err != nil {
		panic(err)
	}
	wires, err := parse(fileContents)
	if err != nil {
		panic(err)
	}
	part1 := solvePart1(wires)
	part2 := solvePart2(wires)
	/*
		Day 3
		Part 1: 1211
		Part 2: 101386
	*/
	fmt.Println("Day 3")
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
