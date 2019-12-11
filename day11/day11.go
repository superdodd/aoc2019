package main

import (
	"fmt"
	"github.com/superdodd/aoc2019/common/intcode"
)

type Point struct {
	X, Y int
}

var day11_input = []int64{
	3, 8, 1005, 8, 345, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8,
	10, 4, 10, 102, 1, 8, 28, 1006, 0, 94, 2, 106, 5, 10, 1, 1109, 12, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10,
	1008, 8, 1, 10, 4, 10, 101, 0, 8, 62, 1, 103, 6, 10, 1, 108, 12, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10,
	1008, 8, 0, 10, 4, 10, 102, 1, 8, 92, 2, 104, 18, 10, 2, 1109, 2, 10, 2, 1007, 5, 10, 1, 7, 4, 10, 3, 8, 102, -1, 8,
	10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 129, 2, 1004, 15, 10, 2, 1103, 15, 10, 2, 1009, 6, 10,
	3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 164, 2, 1109, 14, 10, 1, 1107, 18,
	10, 1, 1109, 13, 10, 1, 1107, 11, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0,
	201, 2, 104, 20, 10, 1, 107, 8, 10, 1, 1007, 5, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4,
	10, 101, 0, 8, 236, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 257, 3, 8, 102,
	-1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 279, 1, 107, 0, 10, 1, 107, 16, 10, 1006, 0, 24,
	1, 101, 3, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 316, 2, 1108, 15, 10,
	2, 4, 11, 10, 101, 1, 9, 9, 1007, 9, 934, 10, 1005, 10, 15, 99, 109, 667, 104, 0, 104, 1, 21101, 0, 936995730328, 1,
	21102, 362, 1, 0, 1105, 1, 466, 21102, 1, 838210728716, 1, 21101, 373, 0, 0, 1105, 1, 466, 3, 10, 104, 0, 104, 1, 3,
	10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102,
	1, 235350789351, 1, 21101, 0, 420, 0, 1105, 1, 466, 21102, 29195603035, 1, 1, 21102, 1, 431, 0, 1105, 1, 466, 3, 10,
	104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 825016079204, 1, 21101, 0, 454, 0, 1105, 1, 466, 21101, 837896786700,
	0, 1, 21102, 1, 465, 0, 1106, 0, 466, 99, 109, 2, 21201, -1, 0, 1, 21101, 0, 40, 2, 21102, 1, 497, 3, 21101, 0, 487, 0,
	1105, 1, 530, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 492, 493, 508, 4, 0, 1001, 492, 1, 492,
	108, 4, 492, 10, 1006, 10, 524, 1101, 0, 0, 492, 109, -2, 2105, 1, 0, 0, 109, 4, 2102, 1, -1, 529, 1207, -3, 0, 10, 1006,
	10, 547, 21102, 1, 0, -3, 21201, -3, 0, 1, 22102, 1, -2, 2, 21101, 1, 0, 3, 21102, 1, 566, 0, 1105, 1, 571, 109, -4, 2106,
	0, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 594, 2207, -4, -2, 10, 1006, 10, 594, 21201, -4, 0, -4, 1106, 0, 662, 21201,
	-4, 0, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 613, 0, 0, 1105, 1, 571, 22101, 0, 1, -4, 21101, 0, 1, -1, 2207,
	-4, -2, 10, 1006, 10, 632, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 654, 22101, 0, -1, 1, 21102,
	654, 1, 0, 105, 1, 529, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0,
}

func paint(hull map[Point]int64) {
	ic := intcode.NewIntcode(day11_input...)
	ic.InputChan = make(chan int64)
	ic.OutputChan = make(chan int64)
	// Start the program in the background
	done := make(chan struct{})
	go func() {
		ic.MustRun()
		close(done)
	}()
	robotDir := 0
	robotLoc := Point{0, 0}
	for done != nil {
		select {
		case ic.InputChan <- hull[robotLoc]:
			hull[robotLoc] = <-ic.OutputChan
			turn := <-ic.OutputChan
			if turn == 0 { // LEFT
				robotDir = (robotDir + 3) % 4
			} else { // RIGHT
				robotDir = (robotDir + 1) % 4
			}
			switch robotDir {
			case 0: // UP
				robotLoc.Y--
			case 1: // RIGHT
				robotLoc.X++
			case 2: // DOWN
				robotLoc.Y++
			case 3: // LEFT
				robotLoc.X--
			}
		case <-done:
			// Program finished
			done = nil
		}
	}
}

func main() {
	hull := map[Point]int64{}
	paint(hull)
	fmt.Println("Day 11: ")
	fmt.Println("Part 1: ", len(hull))
	hull = map[Point]int64{
		{0, 0}: 1,
	}
	paint(hull)
	fmt.Println("Part 2:")
	var maxx, minx, maxy, miny int
	for p := range hull {
		if p.X < minx {
			minx = p.X
		}
		if p.X > maxx {
			maxx = p.X
		}
		if p.Y < miny {
			miny = p.Y
		}
		if p.Y > maxy {
			maxy = p.Y
		}
	}
	img := make([][]bool, maxy-miny+1)
	for i := range img {
		img[i] = make([]bool, maxx-minx+1)
	}
	for p, color := range hull {
		img[p.Y-miny][p.X-minx] = color == 1
	}
	for y := range img {
		for x := range img[y] {
			if img[y][x] {
				fmt.Print("XX")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
