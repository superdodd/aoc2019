package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solvePart1(t *testing.T) {
	pos := []int{
		-1, 0, 2,
		2, -10, -7,
		4, -8, 8,
		3, 5, -1,
	}
	vel := make([]int, len(pos))

	step(pos, vel)
	assert.Equal(t, []int{2, -1, 1, 3, -7, -4, 1, -7, 5, 2, 2, 0}, pos, "pos")
	assert.Equal(t, []int{3, -1, -1, 1, 3, 3, -3, 1, -3, -1, -3, 1}, vel, "vel")

	step(pos, vel)
	assert.Equal(t, []int{5, -3, -1, 1, -2, 2, 1, -4, -1, 1, -4, 2}, pos, "pos2")
	assert.Equal(t, []int{3, -2, -2, -2, 5, 6, 0, 3, -6, -1, -6, 2}, vel, "vel2")

	for i := 2; i < 10; i++ {
		step(pos, vel)
	}
	assert.Equal(t, []int{2, 1, -3, 1, -8, 0, 3, -6, 1, 2, 0, 4}, pos, "pos10")
	assert.Equal(t, []int{-3, -2, 1, -1, 1, 3, 3, 2, -3, 1, -1, -1}, vel, "vel10")
	assert.Equal(t, 179, totalEnergy(pos, vel))

	pos = []int{-8, -10, 0, 5, 5, 10, 2, -7, 3, 9, -8, -3}
	vel = make([]int, len(pos))
	for i := 0; i < 100; i++ {
		step(pos, vel)
	}
	assert.Equal(t, []int{8, -12, -9, 13, 16, -3, -29, -11, -1, 16, -13, 23}, pos, "pos100")
	assert.Equal(t, []int{-7, 3, 0, 3, -11, -5, -3, 7, 4, 7, 1, 1}, vel, "vel100")
	assert.Equal(t, 1940, totalEnergy(pos, vel))
}
