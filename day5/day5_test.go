package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solvePart1(t *testing.T) {
	assert.Equal(t, int64(8332629), solvePart1())
}

func Test_solvePart2(t *testing.T) {
	assert.Equal(t, int64(8805067), solvePart2())
}
