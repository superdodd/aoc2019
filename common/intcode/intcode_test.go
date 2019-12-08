package intcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func runOpcode(instr []int, inputs ...int) (ret *Intcode, err error) {
	ret = NewIntcode(instr...)
	ret.Inputs = append([]int(nil), inputs...)
	err = ret.Step()
	return
}

func TestOpcode_Add(t *testing.T) {
	tests := []struct {
		prog []int
		val  int
	}{
		{[]int{1, 0, 3, 5, 99, 0}, 6},
		{[]int{101, 0, 3, 5, 99, 0}, 5},
		{[]int{1001, 0, 3, 5, 99, 0}, 1004},
		{[]int{1101, 0, 3, 5, 99, 0}, 3},
		{[]int{1101, -1, -2, 5, 99, 0}, -3},
		{[]int{11101, 0, 3, 5, 99, 0}, 3},
		{[]int{10101, 0, 3, 5, 99, 0}, 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("add_%d", i), func(t *testing.T) {
			ic, err := runOpcode(tt.prog)
			assert.Nil(t, err)
			assert.Equal(t, tt.val, ic.Program[5])
			assert.Equal(t, 4, ic.Pc)
			assert.Empty(t, ic.Outputs)
		})
	}
}

func TestOpcode_Mult(t *testing.T) {
	tests := []struct {
		prog []int
		val  int
	}{
		{[]int{2, 0, 3, 5, 99, -1}, 2 * 5},
		{[]int{102, 3, 3, 5, 99, -1}, 3 * 5},
		{[]int{1002, 0, 3, 5, 99, -1}, 1002 * 3},
		{[]int{1102, 3, 3, 5, 99, -1}, 3 * 3},
		{[]int{1102, 1, -1, 5, 99, -1}, 1 * -1},
		{[]int{11102, 2, 3, 5, 99, -1}, 2 * 3},
		{[]int{10102, 2, 3, 5, 99, -1}, 2 * 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("mul_%d", i), func(t *testing.T) {
			ic, err := runOpcode(tt.prog)
			assert.Nil(t, err)
			assert.Equal(t, tt.val, ic.Program[5])
			assert.Equal(t, 4, ic.Pc)
			assert.Empty(t, ic.Outputs)
		})
	}
}

func TestOpcode_Input(t *testing.T) {
	ic, err := runOpcode([]int{3, 3, 99, 0}, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ic.CurrentInput)
	assert.Empty(t, ic.Outputs)
	assert.Equal(t, 2, ic.Pc)
}

func TestOpcode_Output(t *testing.T) {
	ic, err := runOpcode([]int{4, 3, 99, 1})
	assert.Nil(t, err)
	assert.Equal(t, 0, ic.CurrentInput)
	assert.Equal(t, []int{1}, ic.Outputs)
	assert.Equal(t, 2, ic.Pc)
	assert.Equal(t, ic.originalProgram, ic.Program)

	ic, err = runOpcode([]int{104, 3, 99, 1})
	assert.Nil(t, err)
	assert.Equal(t, 0, ic.CurrentInput)
	assert.Equal(t, []int{3}, ic.Outputs)
	assert.Equal(t, 2, ic.Pc)
	assert.Equal(t, ic.originalProgram, ic.Program)
}

func TestOpcode_JumpIfTrue(t *testing.T) {
	tests := []struct {
		prog []int
		pc   int
	}{
		// True
		{[]int{5, 0, 3, 1}, 1},
		{[]int{105, 2, 0, 0}, 105},
		{[]int{1105, 3, 10, 0}, 10},
		{[]int{1005, 0, 2, 1}, 2},
		// False
		{[]int{5, 3, 0, 0}, 3},
		{[]int{105, 0, 0, 0}, 3},
		{[]int{1105, 0, 1, 0}, 3},
		{[]int{1005, 3, 4, 0, 10}, 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("jumpiftrue_%d", i), func(t *testing.T) {
			ic, err := runOpcode(tt.prog)
			assert.Nil(t, err)
			assert.Equal(t, tt.pc, ic.Pc)
			assert.Equal(t, tt.prog, ic.Program)
		})
	}
}

func TestOpcode_JumpIfFalse(t *testing.T) {
	tests := []struct {
		prog []int
		pc   int
	}{
		// True
		{[]int{6, 0, 3, 1}, 3},
		{[]int{106, 2, 0, 0}, 3},
		{[]int{1106, 3, 10, 0}, 3},
		{[]int{1006, 0, 2, 1}, 3},
		// False
		{[]int{6, 3, 0, 0}, 6},
		{[]int{106, 0, 0, 0}, 106},
		{[]int{1106, 0, 1, 0}, 1},
		{[]int{1006, 3, 4, 0, 10}, 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("jumpiffalse_%d", i), func(t *testing.T) {
			ic, err := runOpcode(tt.prog)
			assert.Nil(t, err)
			assert.Equal(t, tt.pc, ic.Pc)
			assert.Equal(t, tt.prog, ic.Program)
		})
	}
}

func TestOpcode_LessThan(t *testing.T) {
	tests := []struct {
		prog []int
		val  int
	}{
		// True
		{[]int{7, 0, 4, 4, 8}, 1},
		{[]int{107, 4, 4, 4, 6}, 1},
		{[]int{1007, 4, 4, 4, 2}, 1},
		{[]int{1107, 10, 11, 4, 2}, 1},
		// False
		{[]int{7, 4, 2, 4, 6}, 0},
		{[]int{107, 109, 0, 4, 6}, 0},
		{[]int{1007, 4, 10, 4, 11}, 0},
		{[]int{1107, 11, 10, 4, 2}, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("lessthan_%d", i), func(t *testing.T) {
			ic, err := runOpcode(tt.prog)
			assert.Nil(t, err)
			assert.Equal(t, tt.val, ic.Program[4])
			assert.NotEqual(t, tt.prog, ic.Program)
			assert.Equal(t, 4, ic.Pc)
		})
	}
}

func TestOpcode_Equals(t *testing.T) {
	tests := []struct {
		prog []int
		val  int
	}{
		// True
		{[]int{8, 0, 4, 4, 8}, 1},
		{[]int{108, 4, 3, 4, 6}, 1},
		{[]int{1008, 4, 10, 4, 10}, 1},
		{[]int{1108, 10, 10, 4, 2}, 1},
		// False
		{[]int{8, 4, 3, 4, 6}, 0},
		{[]int{108, 109, 0, 4, 6}, 0},
		{[]int{1008, 4, 10, 4, 9}, 0},
		{[]int{1108, 10, 11, 4, 2}, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("lessthan_%d", i), func(t *testing.T) {
			ic, err := runOpcode(tt.prog)
			assert.Nil(t, err)
			assert.Equal(t, tt.val, ic.Program[4])
			assert.NotEqual(t, tt.prog, ic.Program)
			assert.Equal(t, 4, ic.Pc)
		})
	}
}

func Test_Day2Examples(t *testing.T) {
	tests := []struct {
		startProgram []int
		endProgram   []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("d2_examples_%d", i), func(t *testing.T) {
			ic := NewIntcode(test.startProgram...)
			ic.MustRun()
			assert.Equal(t, test.endProgram, ic.Program)
		})
	}
}

func Test_Day5Example1(t *testing.T) {
	for i, test := range []int{-1, 0, 1, 5, 1000, 99} {
		ic := NewIntcode(3, 0, 4, 0, 99)
		t.Run(fmt.Sprintf("d5_e1__%d", i), func(t *testing.T) {
			ic.Reset()
			ic.Inputs = []int{test}
			ic.MustRun()
			assert.Equal(t, []int{test}, ic.Outputs)
		})
	}
}

func Test_Day5Examples(t *testing.T) {
	tests := []struct {
		program []int
		inputs  []int
		outputs []int
	}{
		{ // Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8, 0, -8, 3, 99},
			[]int{1, 0, 0, 0, 0},
		},
		{ // Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
			[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			[]int{8, 0, -8, 3, 99},
			[]int{1, 0, 0, 0, 0},
		},
		{ // Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8, 0, -8, 3, 99},
			[]int{0, 1, 1, 1, 0},
		},
		{ // Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
			[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			[]int{8, 0, -8, 3, 99},
			[]int{0, 1, 1, 1, 0},
		},
		// Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero:
		{ // Position mode
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{0, 1, -1, 99},
			[]int{0, 1, 1, 1},
		},
		{ // Immediate mode
			[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			[]int{0, 1, -1, 99},
			[]int{0, 1, 1, 1},
		},
		{
			// uses an input instruction to ask for a single number. The program will then output 999 if the input
			// value is below 8, output 1000 if the input value is equal to 8, or output 1001 if the input value is
			// greater than 8.
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{-1, 0, 1, 8, 9, 99, 1000},
			[]int{999, 999, 999, 1000, 1001, 1001, 1001},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("d5_exapmle_%d", i), func(t *testing.T) {
			ic := NewIntcode(test.program...)
			for input := range test.inputs {
				ic.Reset()
				ic.Inputs = test.inputs[input : input+1]
				ic.MustRun()
				assert.Equal(t, test.outputs[input:input+1], ic.Outputs, "input %d", input)
			}
		})
	}
}

func Test_ChannelInputOutput(t *testing.T) {
	ic := NewIntcode(3, 0, 3, 1, 4, 1, 4, 0, 99)
	inputChan := make(chan int)
	outputChan := make(chan int)
	ic.InputChan = inputChan
	ic.OutputChan = outputChan
	done := make(chan struct{})
	go func() {
		ic.MustRun()
		close(done)
	}()
	inputChan <- 2
	inputChan <- 3
	close(inputChan)
	o, ok := <-outputChan
	assert.True(t, ok)
	assert.Equal(t, 3, o)
	o, ok = <-outputChan
	assert.True(t, ok)
	assert.Equal(t, 2, o)
	timeout := time.After(1 * time.Second)
	select {
	case <-done:
	case <-timeout:
		t.Fatal("timeout!")
	}
}

func Test_ChainedInputOutput(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int, 1)
	ic1 := NewIntcode(1, 0, 1, 0, 3, 0, 1, 0, 2, 0, 4, 0, 99)
	ic1.InputChan = a
	ic1.OutputChan = b
	ic2 := NewIntcode(ic1.Program...)
	ic2.InputChan = b
	ic2.OutputChan = a
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ic1.MustRun()
		wg.Done()
	}()
	go func() {
		ic2.MustRun()
		wg.Done()
	}()
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	a <- 0
	timeout := time.After(1 * time.Second)
	select {
	case <-done:
	case <-timeout:
		t.Fatal("timeout!")
	}

}
