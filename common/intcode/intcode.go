package intcode

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Intcode struct {
	Pc              int   // The current instruction pointer
	Program         []int // The current working memory
	Inputs          []int // The sequence of input values for the program
	CurrentInput    int   // The index of the next input to be consumed
	Outputs         []int // The outputs from running the program, in order
	originalProgram []int // The program as originally loaded
}

func NewFromFile(fileName string) *Intcode {
	ret := &Intcode{}
	ret.MustLoadFile(fileName)
	return ret
}

func NewFromString(contents string) *Intcode {
	ret := &Intcode{}
	ret.MustLoad(contents)
	return ret
}

func NewIntcode(mem ...int) *Intcode {
	return &Intcode{
		Program:         append([]int(nil), mem...),
		originalProgram: append([]int(nil), mem...),
	}
}

// Reset the program memory to the initial input.  Returns the previous memory state.
func (ic *Intcode) Reset() []int {
	ret := ic.Program
	ic.Program = append([]int(nil), ic.originalProgram...)
	ic.Outputs = nil
	ic.Pc = 0
	ic.CurrentInput = 0
	return ret
}

// Loads a program from a file.
func (ic *Intcode) LoadFile(fileName string) error {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return ic.Load(string(fileContents))
}

// Loads a program from a file or panics in the event of an error.
func (ic *Intcode) MustLoadFile(fileName string) {
	if err := ic.LoadFile(fileName); err != nil {
		panic(err)
	}
}

// Loads a program from a string.
func (ic *Intcode) Load(contents string) error {
	ic.originalProgram = nil
	for i, num := range strings.Split(contents, ",") {
		ival, err := strconv.Atoi(num)
		if err != nil {
			return fmt.Errorf("loc %d: %v", i, err)
		}
		ic.originalProgram = append(ic.originalProgram, ival)
	}
	ic.Reset()
	return nil
}

// Loads a program from a string or panics in the event of an error.
func (ic *Intcode) MustLoad(contents string) {
	if err := ic.Load(contents); err != nil {
		panic(err)
	}
}

// Returns the relevant value (accounting for immediate vs position) for the given parameter number
// for the currently active instruction.
func (ic *Intcode) paramVal(paramNumber int) int {
	if (ic.Program[ic.Pc]/int(math.Pow10(paramNumber+1)))%10 == 0 {
		return ic.Program[ic.Program[ic.Pc+paramNumber]]
	}
	return ic.Program[ic.Pc+paramNumber]
}

// Executes the currently resident program in memory.  Starts executing at the current value of Pc.
func (ic *Intcode) Run() error {
	for ic.Program[ic.Pc]%100 != 99 {
		if err := ic.Step(); err != nil {
			return err
		}
	}
	return nil
}

func (ic *Intcode) makeErr(msg string, args ...interface{}) error {
	return fmt.Errorf(msg+" at pc=%d", append(args, ic.Pc)...)
}

// Executes a single instruction and advances the program counter.
func (ic *Intcode) Step() error {
	switch ic.Program[ic.Pc] % 100 {
	case 1: // Add
		ic.Program[ic.Program[ic.Pc+3]] = ic.paramVal(1) + ic.paramVal(2)
		ic.Pc += 4
	case 2: // Multiply
		ic.Program[ic.Program[ic.Pc+3]] = ic.paramVal(1) * ic.paramVal(2)
		ic.Pc += 4
	case 3: // Input
		if ic.CurrentInput >= len(ic.Inputs) {
			return ic.makeErr("no more inputs")
		}
		ic.Program[ic.Program[ic.Pc+1]] = ic.Inputs[ic.CurrentInput]
		ic.CurrentInput++
		ic.Pc += 2
	case 4: // Output
		ic.Outputs = append(ic.Outputs, ic.paramVal(1))
		ic.Pc += 2
	case 5: // Jump if true
		if ic.paramVal(1) != 0 {
			ic.Pc = ic.paramVal(2)
		} else {
			ic.Pc += 3
		}
	case 6: // Jump if false
		if ic.paramVal(1) == 0 {
			ic.Pc = ic.paramVal(2)
		} else {
			ic.Pc += 3
		}
	case 7: // Less than
		if ic.paramVal(1) < ic.paramVal(2) {
			ic.Program[ic.Program[ic.Pc+3]] = 1
		} else {
			ic.Program[ic.Program[ic.Pc+3]] = 0
		}
		ic.Pc += 4
	case 8: // Equals
		if ic.paramVal(1) == ic.paramVal(2) {
			ic.Program[ic.Program[ic.Pc+3]] = 1
		} else {
			ic.Program[ic.Program[ic.Pc+3]] = 0
		}
		ic.Pc += 4
	default: // Error
		return ic.makeErr("unexpected opcode: %d", ic.Program[ic.Pc])
	}
	return nil
}

// Runs a program, panic on errors
func (ic *Intcode) MustRun() {
	if err := ic.Run(); err != nil {
		panic(err)
	}
}

func (ic *Intcode) InitializeProgram(mem ...int) {
	ic.originalProgram = append([]int(nil), mem...)
	ic.Reset()
}
