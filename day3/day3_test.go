package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCoord_distance(t *testing.T) {
	tests := []struct {
		this  Coord
		other Coord
		want  int
	}{
		{Coord{0, 0}, Coord{1, 1}, 2},
		{Coord{1, 1}, Coord{1, 1}, 0},
		{Coord{1, 1}, Coord{0, 0}, 2},
		{Coord{1, 3}, Coord{0, 0}, 4},
		{Coord{3, 2}, Coord{0, 0}, 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("distance_%d", i), func(t *testing.T) {
			if got := tt.this.Distance(tt.other); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		input   string
		want    [][]Coord
		wantErr bool
	}{
		{
			"R1,U2,L2,D2",
			[][]Coord{
				{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {0, 2}, {-1, 2}, {-1, 1}, {-1, 0}},
			}, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("parse_%d", i), func(t *testing.T) {
			got, err := parse([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	tests := []struct {
		input string
		want int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("part1_%d", i), func(t *testing.T) {
			wires, err := parse([]byte(tt.input))
			if err != nil {
				t.Errorf("solvePart1() err %v", err)
			}
			if got := solvePart1(wires); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	tests := []struct {
		input string
		want int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("part1_%d", i), func(t *testing.T) {
			wires, err := parse([]byte(tt.input))
			if err != nil {
				t.Errorf("solvePart2() err %v", err)
			}
			if got := solvePart2(wires); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}