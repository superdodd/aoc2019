package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/superdodd/aoc2019/common/intcode"
	"testing"
)

func Test_BOOST(t *testing.T) {
	ic := intcode.NewIntcode(day9_input...)
	ic.Inputs = []int64{1}
	ic.MustRun()
	assert.Equal(t, 1, len(ic.Outputs))
	assert.Equal(t, int64(2518058886), ic.Outputs[0])
}

func Test_BOOST2(t *testing.T) {
	ic := intcode.NewIntcode(day9_input...)
	ic.Inputs = []int64{2}
	ic.MustRun()
	assert.Equal(t, 1, len(ic.Outputs))
	assert.Equal(t, int64(44292), ic.Outputs[0])
}
