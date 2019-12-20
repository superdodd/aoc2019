package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlawedFrequencyTransform(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{"12345678", []int{4, 8, 2, 2, 6, 1, 5, 8}},
		{"48226158", []int{3, 4, 0, 4, 0, 4, 3, 8}},
		{"34040438", []int{0, 3, 4, 1, 5, 5, 1, 8}},
		{"03415518", []int{0, 1, 0, 2, 9, 4, 9, 8}},
	}
	for tti, tt := range tests {
		t.Run(fmt.Sprintf("fft_%d", tti), func(t *testing.T) {
			if got := FlawedFrequencyTransform(Parse(tt.in)); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("FlawedFrequencyTransform() = %v, want %v", got, tt.out)
			}
		})
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	for _, reps := range []int{1, 10, 100, 500} {
		b.Run(fmt.Sprintf("reps_%d", reps), func(b *testing.B) {
			input := Parse(day16_input)
			extended_input := make([]int, len(input)*reps)
			for r := 0; r < reps; r++ {
				copy(extended_input[len(input)*r:len(input)*(r+1)], input)
			}
			for i := 0; i < b.N; i++ {
				SolvePart1(extended_input)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"03036732577212944063491565474664", "84462026"},
		{"02935109699940807407585447034323", "78725270"},
		{"03081770884921959731165446850517", "53553731"},
	}
	for tti, tt := range tests {
		t.Run(fmt.Sprintf("part2_%d", tti), func(t *testing.T) {
			got := SolvePart2(Parse(tt.input))
			if got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
