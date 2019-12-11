package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var day10_input = "" +
	".#......##.#..#.......#####...#..\n" +
	"...#.....##......###....#.##.....\n" +
	"..#...#....#....#............###.\n" +
	".....#......#.##......#.#..###.#.\n" +
	"#.#..........##.#.#...#.##.#.#.#.\n" +
	"..#.##.#...#.......#..##.......##\n" +
	"..#....#.....#..##.#..####.#.....\n" +
	"#.............#..#.........#.#...\n" +
	"........#.##..#..#..#.#.....#.#..\n" +
	".........#...#..##......###.....#\n" +
	"##.#.###..#..#.#.....#.........#.\n" +
	".#.###.##..##......#####..#..##..\n" +
	".........#.......#.#......#......\n" +
	"..#...#...#...#.#....###.#.......\n" +
	"#..#.#....#...#.......#..#.#.##..\n" +
	"#.....##...#.###..#..#......#..##\n" +
	"...........#...#......#..#....#..\n" +
	"#.#.#......#....#..#.....##....##\n" +
	"..###...#.#.##..#...#.....#...#.#\n" +
	".......#..##.#..#.............##.\n" +
	"..###........##.#................\n" +
	"###.#..#...#......###.#........#.\n" +
	".......#....#.#.#..#..#....#..#..\n" +
	".#...#..#...#......#....#.#..#...\n" +
	"#.#.........#.....#....#.#.#.....\n" +
	".#....#......##.##....#........#.\n" +
	"....#..#..#...#..##.#.#......#.#.\n" +
	"..###.##.#.....#....#.#......#...\n" +
	"#.##...#............#..#.....#..#\n" +
	".#....##....##...#......#........\n" +
	"...#...##...#.......#....##.#....\n" +
	".#....#.#...#.#...##....#..##.#.#\n" +
	".#.#....##.......#.....##.##.#.##"

type Point struct {
	X, Y     int
	Angle    float64
	Distance float64
	Order    int
}

func parseField(input string) []Point {
	var ret []Point
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				ret = append(ret, Point{X: x, Y: y})
			}
		}
	}
	return ret
}

func countVisible(field []Point, origin Point) int {
	var candidates []Point
	for _, p := range field {
		if p.X != origin.X || p.Y != origin.Y {
			candidates = append(candidates, p)
		}
	}
	// Sort the candidates by distance from our current origin.
	sort.Slice(candidates, func(i, j int) bool {
		return math.Abs(float64(candidates[i].X-origin.X))+math.Abs(float64(candidates[i].Y-origin.Y)) < math.Abs(float64(candidates[j].X-origin.X))+math.Abs(float64(candidates[j].Y-origin.Y))
	})

	visibleCount := 0
	// For each candidate, check to see if there are any closer asteroids with the same slope from the origin.
candidate:
	for i, c := range candidates {
		for j, o := range candidates {
			if j >= i {
				break
			}
			if (c.X == origin.X && o.X == origin.X) ||
				(float64(c.Y-origin.Y)/float64(c.X-origin.X) == float64(o.Y-origin.Y)/float64(o.X-origin.X)) {
				// This candidate is aligned with some closer one; check that they are in the same quadrant.
				// Since candidates are sorted by distance from origin, o.Y should be strictly between c.Y and origin.Y
				// if o is actually in the way.
				if ((c.Y <= o.Y && o.Y <= origin.Y) || (c.Y >= o.Y && o.Y >= origin.Y)) &&
					((c.X <= o.X && o.X <= origin.X) || (c.X >= o.X && o.X >= origin.X)) {
					continue candidate
				}
			}
		}
		visibleCount++
	}
	return visibleCount
}

func solvePart1(field []Point) (x, y, maxSeen int) {
	for _, o := range field {
		s := countVisible(field, o)
		if s > maxSeen {
			maxSeen, x, y = s, o.X, o.Y
		}
	}
	return
}

func solvePart2(field []Point, base Point) (x, y int) {
	var candidates []Point
	for _, p := range field {
		if p.X != base.X || p.Y != base.Y {
			p.Angle = math.Atan2(float64(p.X-base.X), float64(base.Y-p.Y))
			if p.Angle < 0 {
				p.Angle += 2 * math.Pi
			}
			p.Distance = math.Abs(float64(p.X-base.X)) + math.Abs(float64(p.Y-base.Y))
			candidates = append(candidates, p)
		}
	}
	// First sort by angle and distance...
	sort.Slice(candidates, func(i, j int) bool {
		a := candidates[i].Angle - candidates[j].Angle
		if a != 0 {
			return a < 0
		}
		return candidates[i].Distance < candidates[j].Distance
	})
	// Then increase the order of any asteroids that are at the same angle.
	var maxOrder int
	for i := range candidates {
		if i > 0 && candidates[i-1].Angle == candidates[i].Angle {
			candidates[i].Order = candidates[i-1].Order + 1
			if candidates[i].Order > maxOrder {
				maxOrder = candidates[i].Order
			}
		}
	}

	// Finally, loop through the list, vaporizing in order.
	var vaporizedCount = 0
	for o := 0; o <= maxOrder; o++ {
		for _, p := range candidates {
			if p.Order != o {
				continue
			}
			vaporizedCount++
			if vaporizedCount == 200 {
				x, y = p.X, p.Y
				return
			}
		}
	}
	panic(fmt.Errorf("only vaporized %d", vaporizedCount))
}

func main() {
	asteroidField := parseField(day10_input)
	fmt.Println("Day 10:")
	x, y, c := solvePart1(asteroidField)
	fmt.Printf("Part 1: %d (%d, %d)\n", c, x, y)
	x, y = solvePart2(asteroidField, Point{X: x, Y: y})
	fmt.Println("Part 2: ", 100*x+y)
}
