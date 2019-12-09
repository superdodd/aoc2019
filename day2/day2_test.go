package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/superdodd/aoc2019/common/intcode"
	"testing"
)

func Test_solvePart1(t *testing.T) {
	assert.Equal(t, int64(5434663), solvePart1(intcode.NewIntcode(day2_program...)))
}

func Test_solvePart2(t *testing.T) {
	assert.Equal(t, int64(4559), solvePart2(intcode.NewIntcode(day2_program...)))
}
