package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		startOpcodes []int
		endOpcodes []int
	}{
		{[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
		{[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}},
		{[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			runProgram(test.startOpcodes)
			assert.Equal(t, test.endOpcodes, test.startOpcodes)
		})
	}
}
