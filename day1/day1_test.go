package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestSolve(t *testing.T) {
	tests := []struct {
		input int
		output int
		output2 int
	}{
		{12, 2, 2},
		{14, 2, 2},
		{1969, 654, 966},
		{100756, 33583, 50346},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("input_%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.output, solve([]int{test.input}))
			assert.Equal(t, test.output2, solvePart2([]int{test.input}))
		})
	}
}
