package main

import "fmt"

var day16_input = "597235178986903423360856190279211112600006674170525294338940926497796855575579963830857089032415354367" +
	"86723718804155370155263736632861535632645335233170435646844328735934063129720822438983948765830873108060969395372667" +
	"94408120102015412673656521245540358256581403756833210604333665797290629730699372771473006102932115398439065894901382" +
	"19183523415036297055876667796810133580533129907094231561102918357941790564329585377968552877342171256157001999289155" +
	"24410743382078079059706420865085147514027374485354815106354367548002650415494525590292210440827027951624280115914909" +
	"910917047084328588833201558964370296841789611989343040407348115608623432403085634084"

func Parse(input string) []int {
	var ret = make([]int, len(input))
	for i, r := range input {
		ret[i] = int(r - '0')
		if ret[i] < 0 || ret[i] > 9 {
			panic("invalid parse")
		}
	}
	return ret
}

func FlawedFrequencyTransform(input []int) []int {
	pattern := []int{0, 1, 0, -1}
	out := make([]int, len(input))
	for o := range out {
		for i, v := range input {
			out[o] += v * pattern[(((i+1)/(o+1))%len(pattern))]
		}
		out[o] %= 10
		if out[o] < 0 {
			out[o] = -out[o]
		}
	}
	return out
}

func SolvePart1(input []int) string {
	out := append([]int(nil), input...)
	for i := 0; i < 100; i++ {
		out = FlawedFrequencyTransform(out)
	}
	var ret string
	for i := 0; i < 8; i++ {
		ret += string('0' + out[i])
	}
	return ret
}

func FlawedFrequencyTransformSecondHalfOfInput(input []int) []int {
	// We observe that given an input length of L, all output positions o >= L/2 are computed by summing up each digit
	// from o to the end of the input:
	// If L = 10, then output position 5 is the sum of input digits 5-10, position 6 is the sum of the input digits 6-10,
	// and so on.  So it is easy to calculate digits for the last half of the output by working *backwards* from the end
	// of the input, carrying forward the total from previous digits and simply adding the next digit to the total as
	// we work backwards.  We can stop as soon as we get to the earliest output position we're interested in.
	output := make([]int, len(input))
	copy(output, input)
	for o := len(input) - 2; o >= 0; o-- {
		output[o] += output[o+1]
	}
	for o := range output {
		output[o] %= 10
	}
	return output
}

func SolvePart2(input []int) string {
	var targetPosition int
	for i := 0; i < 7; i++ {
		targetPosition = 10*targetPosition + input[i]
	}
	targetPosition--
	totalInputLen := len(input) * 10000
	extendedInput := make([]int, totalInputLen-targetPosition-1)
	n := copy(extendedInput, input[(targetPosition%len(input))+1:])
	for n < len(extendedInput) {
		c := copy(extendedInput[n:], input)
		if c < len(input) {
			panic("invalid copy")
		}
		n += c
	}
	for i := 0; i < 100; i++ {
		extendedInput = FlawedFrequencyTransformSecondHalfOfInput(extendedInput)
	}
	var ret string
	for i := 0; i < 8; i++ {
		ret += string('0' + extendedInput[i])
	}
	return ret
}

func main() {
	fmt.Println("Day 16:")
	fmt.Println("Part 1: ", SolvePart1(Parse(day16_input)))
	fmt.Println("Part 2: ", SolvePart2(Parse(day16_input)))
}
