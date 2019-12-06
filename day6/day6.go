package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type OrbitNode struct {
	Label  string
	Parent *OrbitNode
}

func Parse(fileContents string) map[string]*OrbitNode {
	ret := map[string]*OrbitNode{}
	for lineno, line := range strings.Split(fileContents, "\n") {
		labels := strings.Split(line, ")")
		if len(labels) != 2 {
			panic(fmt.Errorf("unrecognized format at line %d: %s", lineno, line))
		}
		// Parent object being orbited
		parent, ok := ret[labels[0]]
		if !ok {
			parent = &OrbitNode{
				Label:  labels[0],
				Parent: nil,
			}
			ret[labels[0]] = parent
		}
		// Satellite
		satellite, ok := ret[labels[1]]
		if !ok {
			satellite = &OrbitNode{
				Label:  labels[1],
				Parent: parent,
			}
			ret[labels[1]] = satellite
		} else {
			if satellite.Parent == nil {
				satellite.Parent = parent
			} else if satellite.Parent.Label != labels[1] {
				panic(fmt.Errorf("multiple parents at line %d: %s", lineno, line))
			}
		}
	}
	return ret
}

func CountTotalOrbits(orbits map[string]*OrbitNode) int {
	ret := 0
	for _, o := range orbits {
		for i := o; i.Parent != nil; i = i.Parent {
			ret++
		}
	}
	return ret
}

func PathToCom(orbits map[string]*OrbitNode, label string) []string {
	var ret []string
	for p := orbits[label]; p.Parent != nil; p = p.Parent {
		// Push each element onto the front
		ret = append([]string{p.Parent.Label}, ret...)
	}
	return ret
}

func PathBetweenOrbits(orbits map[string]*OrbitNode, you, san string) []string {
	youList := PathToCom(orbits, you)
	fmt.Printf("PATH_BETWEEN_ORBITS: YOU -> COM: %d\n", len(youList))
	sanList := PathToCom(orbits, san)
	fmt.Printf("PATH_BETWEEN_ORBITS: SAN -> COM: %d\n", len(sanList))
	var i int
	for i = 0; i < len(youList) && i < len(sanList); i++ {
		if youList[i] != sanList[i] {
			// Previous element was the last one in common
			i--
			fmt.Printf("PATH_BETWEEN_ORBITS common index = %d: %s\n", i, youList[i])
			break
		}
	}
	if i < 0 {
		panic(fmt.Errorf("invalid orbits! No common path found: %+v, %+v", youList, sanList))
	}
	var ret []string
	for j := len(youList) - 1; j > i; j-- {
		ret = append(ret, youList[j])
	}
	for j := i; j < len(sanList); j++ {
		ret = append(ret, sanList[j])
	}
	return ret
}

func main() {
	fileContents, err := ioutil.ReadFile("src/github.com/superdodd/aoc2019/day6/day6_input.txt")
	if err != nil {
		panic(err)
	}
	nodes := Parse(string(fileContents))
	fmt.Println("Day 6:")
	fmt.Println("Part 1: ", CountTotalOrbits(nodes))
	fmt.Println("Part 2: ", len(PathBetweenOrbits(nodes, "YOU", "SAN"))-1)
}
