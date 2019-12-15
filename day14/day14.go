package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const day14_input = `
4 SRWZ, 3 ZGSFW, 1 HVJVQ, 1 RWDSX, 12 BDHX, 1 GDPKF, 23 WFPSM, 1 MPKC => 6 VCWNW
3 BXVJK, 3 WTPN => 4 GRQC
5 KWFD => 9 NMZND
1 DNZQ, 5 CDSP => 3 PFDBV
4 VSPSC, 34 MPKC, 9 DFNVL => 9 PZWSP
5 NTXHM => 9 DBKN
4 JNSP, 4 TCKR, 7 PZWSP => 7 DLHG
12 CNBS, 3 FNPC => 2 SRWZ
3 RWDSX, 4 NHSTB, 2 JNSP => 8 TCKR
24 PGHF, 1 NMZND => 3 RWDSX
1 DLHG => 9 QSVN
6 HVJVQ => 2 QSNCW
4 CHDTJ => 9 FDVNC
1 HBXF, 1 RWDSX => 7 BWSPN
2 ZGSFW, 1 KWFD => 8 JNSP
2 BWSPN, 7 GDPKF, 1 BXVJK => 6 FVQM
2 MHBH => 6 FNPC
2 WTPN, 15 GRQC => 3 ZGSFW
9 LXMLX => 6 CLZT
5 DFNVL, 1 KHCQ => 4 MHLBR
21 CNTFK, 3 XHST => 9 CHDTJ
1 CNTFK => 7 MHBH
1 GMQDW, 34 GDPKF, 2 ZDGPL, 1 HVJVQ, 13 QSVN, 1 QSNCW, 1 BXVJK => 2 SGLGN
1 BMVRK, 1 XHST => 8 XHLNT
23 CXKN => 1 BDKN
121 ORE => 9 XHST
4 NTXHM, 4 FNPC, 15 VCMVN => 8 MPKC
2 ZDGPL, 7 JNSP, 3 FJVMD => 4 GMQDW
1 LXMLX, 2 BWSPN => 2 DNZQ
6 WTPN => 9 KCMH
20 CDSP => 2 VSPSC
2 QSNCW, 1 BDHX, 3 HBXF, 8 PFDBV, 17 ZDGPL, 1 MHLBR, 9 ZGSFW => 8 FDWSG
2 VSFTG, 2 DLHG => 9 BDHX
174 ORE => 5 BMVRK
2 BMVRK => 2 KWFD
3 WTPN, 9 TVJPG => 9 CDSP
191 ORE => 2 CNTFK
9 FDVNC, 1 MHBH => 8 NTXHM
3 NHSTB, 2 BXVJK, 1 JNSP => 1 WFPSM
7 FJVMD => 9 CXKN
3 GDPKF, 10 QSNCW => 7 ZDGPL
7 LPXM, 11 VSPSC => 1 LXMLX
6 RWDSX, 2 NMZND, 1 MPKC => 1 KHCQ
6 RWDSX => 4 QMJK
15 MHBH, 28 DBKN, 12 CNBS => 4 PGHF
20 NMZND, 1 PGHF, 1 BXVJK => 2 LPXM
1 CDSP, 17 BXVJK => 5 NHSTB
12 HVJVQ => 3 VSFTG
2 PGHF, 3 VCMVN, 2 NHSTB => 1 DFNVL
5 FNPC => 9 HBXF
3 DPRL => 4 FJVMD
1 KWFD, 1 TVJPG => 8 VCMVN
1 FDWSG, 1 VCWNW, 4 BDKN, 14 FDVNC, 1 CLZT, 62 SGLGN, 5 QMJK, 26 ZDGPL, 60 KCMH, 32 FVQM, 15 SRWZ => 1 FUEL
3 XHLNT => 8 TVJPG
5 HBXF => 2 HVJVQ
3 CHDTJ, 15 KWFD => 9 WTPN
7 CNTFK => 7 CNBS
1 CNBS => 2 JPDF
5 JNSP => 8 DPRL
11 NTXHM => 8 GDPKF
10 JPDF => 9 BXVJK`

type Ingredient struct {
	Name         string
	Makes        []*Ingredient
	RecipeOutput int
	RecipeInput  map[*Ingredient]int
}

func Parse(in string) map[string]*Ingredient {
	allIngredients := map[string]*Ingredient{}
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			continue
		}
		matches := regexp.MustCompile(`\d+ [A-Z]+`).FindAllString(line, 20)
		names := make([]string, len(matches))
		counts := map[string]int{}
		for i := range matches {
			fields := strings.Split(matches[i], " ")
			names[i] = fields[1]
			count, err := strconv.Atoi(fields[0])
			if err != nil {
				panic(err)
			}
			counts[fields[1]] = count
			ingredient, ok := allIngredients[fields[1]]
			if !ok {
				ingredient = &Ingredient{
					Name:        fields[1],
					RecipeInput: map[*Ingredient]int{},
				}
				allIngredients[fields[1]] = ingredient
			}
		}
		output := allIngredients[names[len(names)-1]]
		output.RecipeOutput = counts[output.Name]
		for i := 0; i < len(matches)-1; i++ {
			input := allIngredients[names[i]]
			input.Makes = append(input.Makes, output)
			output.RecipeInput[input] = counts[names[i]]
		}
	}
	return allIngredients
}

func (i *Ingredient) neededToMake(o *Ingredient) bool {
	for _, x := range i.Makes {
		if x == o || x.neededToMake(o) {
			return true
		}
	}
	return false
}

func SolvePart1(ingredients map[string]*Ingredient) int {
	return int(analyzeOreRequirement(ingredients, true))
}

func SolvePart2(ingredients map[string]*Ingredient) int {
	exactOrePerFuel := analyzeOreRequirement(ingredients, false)
	return int(1000000000000 / exactOrePerFuel)
}

func analyzeOreRequirement(ingredients map[string]*Ingredient, wholeBatchesOnly bool) float64 {
	// First sort in reverse dependency order.  We want to ensure that we only
	// attempt to decompose an ingredient into its components once.
	// Not sure how efficient this sort is, but at least it's obvious.
	topoSort := make([]*Ingredient, 0, len(ingredients))
	unsortedIngredients := make(map[string]*Ingredient, len(ingredients))
	for k, v := range ingredients {
		unsortedIngredients[k] = v
	}
	for len(unsortedIngredients) > 0 {
		var newSort *Ingredient
	u:
		for _, i := range unsortedIngredients {
			// If everything directly made by this recipe is already in the sorted list, add this one to the head of
			// the sorted list and delete it from the map.
			for _, m := range i.Makes {
				if _, ok := unsortedIngredients[m.Name]; ok {
					continue u
				}
			}
			newSort = i
		}
		if newSort == nil {
			panic("invalid sort")
		}
		delete(unsortedIngredients, newSort.Name)
		topoSort = append(topoSort, newSort)
	}

	outputsRequired := map[*Ingredient]float64{ingredients["FUEL"]: 1}
	var oreRequired float64
	for len(outputsRequired) > 0 {
		var out *Ingredient
		var requiredCount float64
		for _, i := range topoSort {
			var ok bool
			if requiredCount, ok = outputsRequired[i]; ok {
				out = i
				delete(outputsRequired, i)
				break
			}
		}
		if out == nil {
			panic("Unable to make recipe")
		}

		// Decompose the needed ingredient into its recipe components and add each to the required list.
		batchesRequired := requiredCount / float64(out.RecipeOutput)
		if wholeBatchesOnly {
			batchesRequired = math.Ceil(batchesRequired)
		}
		for ingredient, count := range out.RecipeInput {
			if ingredient.Name == "ORE" {
				oreRequired += batchesRequired * float64(count)
			}
			outputsRequired[ingredient] = outputsRequired[ingredient] + batchesRequired*float64(count)
		}
	}
	return oreRequired
}

func main() {
	ingredients := Parse(day14_input)
	fmt.Println("Day 14:")
	fmt.Println("Part 1: ", SolvePart1(ingredients))
	fmt.Println("Part 2: ", SolvePart2(ingredients))
}
