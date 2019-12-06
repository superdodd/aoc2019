package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	B := &OrbitNode{"B", nil}
	C := &OrbitNode{"C", nil}
	D := &OrbitNode{"D", nil}
	E := &OrbitNode{"E", nil}
	F := &OrbitNode{"F", nil}
	G := &OrbitNode{"G", nil}
	H := &OrbitNode{"H", nil}
	I := &OrbitNode{"I", nil}
	J := &OrbitNode{"J", nil}
	K := &OrbitNode{"K", nil}
	L := &OrbitNode{"L", nil}
	COM := &OrbitNode{"COM", nil}
	B.Parent = COM
	G.Parent = B
	H.Parent = G
	C.Parent = B
	D.Parent = C
	I.Parent = D
	E.Parent = D
	J.Parent = E
	F.Parent = E
	K.Parent = J
	L.Parent = K
	parseResults := map[string]*OrbitNode{
		"B": B, "C": C, "D": D, "E": E, "F": F, "G": G, "H": H, "I": I, "J": J, "K": K, "L": L, "COM": COM,
	}
	assert.Equal(t, parseResults, Parse("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"))
}

func TestCountTotalOrbits(t *testing.T) {
	assert.Equal(t, 42, CountTotalOrbits(Parse("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L")))
}

func TestPathBetweenOrbits(t *testing.T) {
	orbits := Parse("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN")
	assert.Equal(t,
		[]string{"K", "J", "E", "D", "I"},
		PathBetweenOrbits(orbits, "YOU", "SAN"))
}
