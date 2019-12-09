package intcode

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"strings"
	"sync/atomic"
)

const MAX_PROG_MEM = 1 << 16

var serialNum uint64

type Intcode struct {
	Pc              int64           // The current instruction pointer
	RelativeOffset  int64           // The current relative offset
	Program         []int64         // The current working memory
	ExtendedMemory  map[int64]int64 // Extended memory locations beyond initial program space
	Inputs          []int64         // The sequence of input values for the program
	CurrentInput    int             // The index of the next input to be consumed
	InputChan       chan int64      // Read-only channel of input values (may be nil)
	Outputs         []int64         // The outputs from running the program, in order
	OutputChan      chan int64      // Write-only channel of output values (may be nil)
	originalProgram []int64         // The program as originally loaded
	SerialNum       uint64          // A unique identifier for this instance (auto-generated)
}

func NewIntcode(mem ...int64) *Intcode {
	return &Intcode{
		Program:         append([]int64(nil), mem...),
		originalProgram: append([]int64(nil), mem...),
		ExtendedMemory:  map[int64]int64{},
		SerialNum:       atomic.AddUint64(&serialNum, 1),
	}
}

func (ic *Intcode) Copy() *Intcode {
	ret := NewIntcode(ic.Program...)
	ret.Inputs = append([]int64(nil), ic.Inputs...)
	for loc, val := range ic.ExtendedMemory {
		ret.ExtendedMemory[loc] = val
	}
	return ret
}

// Reset the program memory to the initial input.  Returns the previous memory state.
func (ic *Intcode) Reset() []int64 {
	ret := ic.Program
	ic.Program = append([]int64(nil), ic.originalProgram...)
	ic.ExtendedMemory = map[int64]int64{}
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
		ival, ok := new(big.Int).SetString(num, 10)
		if !ok {
			return fmt.Errorf("loc %d: SetString failed", i)
		}
		ic.originalProgram = append(ic.originalProgram, ival.Int64())
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

// Creates a new Intcode program with the given starting memory, whose input channel is connected to the output
// channel of this program via a buffered channel.
func (ic *Intcode) Chain(mem ...int64) *Intcode {
	ret := NewIntcode(mem...)
	if ic == nil {
		// Allow chaining from a nil pointer - just create an input channel
		ret.InputChan = make(chan int64, 1)
	} else {
		if ic.OutputChan == nil {
			ic.OutputChan = make(chan int64, 1)
		}
		ret.InputChan = ic.OutputChan
	}
	return ret
}

// Returns the relevant value (accounting for immediate vs position) for the given parameter number
// for the currently active instruction.
func (ic *Intcode) paramVal(paramNumber int) int64 {
	pVal := ic.readMem(ic.Pc + int64(paramNumber))
	paramMode := (ic.readMem(ic.Pc) / int64(math.Pow10(paramNumber+1))) % 10
	switch paramMode {
	case 0:
		// Positional
		return ic.readMem(pVal)
	case 1:
		// Immediate
		return pVal
	case 2:
		// Relative
		return ic.readMem(pVal + ic.RelativeOffset)
	default:
		panic(fmt.Errorf("unrecognized parameter mode %d for param %d: %d", paramMode, paramNumber, ic.readMem(ic.Pc)))
	}
}

func (ic *Intcode) readMem(loc int64) int64 {
	if loc < int64(len(ic.Program)) {
		return ic.Program[loc]
	}
	if loc < MAX_PROG_MEM {
		// Expand program memory to allow for direct access - aligned to 1024-elements
		newMem := make([]int64, ((loc/1024)+1)*1024)
		copy(newMem, ic.Program)
		ic.Program = newMem
		return ic.Program[loc]
	}
	return ic.ExtendedMemory[loc]
}

func (ic *Intcode) setMem(paramNumber int, val int64) {
	pVal := ic.readMem(ic.Pc + int64(paramNumber))
	paramMode := (ic.readMem(ic.Pc) / int64(math.Pow10(paramNumber+1))) % 10
	var loc int64
	switch paramMode {
	case 0, 1:
		// Positional.  Assume i
		loc = pVal
	case 2:
		// Relative
		loc = pVal + ic.RelativeOffset
	default:
		panic(fmt.Errorf("unrecognized parameter mode %d for param %d: %d", paramMode, paramNumber, ic.readMem(ic.Pc)))
	}
	if loc < int64(len(ic.Program)) {
		// "Direct access"
		ic.Program[int(loc)] = val
		return
	}
	if loc < MAX_PROG_MEM {
		// Expand program memory to allow for direct access - aligned to 1024-elements
		newMem := make([]int64, ((loc/1024)+1)*1024)
		copy(newMem, ic.Program)
		ic.Program = newMem
		ic.Program[int(loc)] = val
		return
	}
	// "Indirect access"
	ic.ExtendedMemory[loc] = val
}

// Executes the currently resident program in memory.  Starts executing at the current value of Pc.
func (ic *Intcode) Run() error {
	for ic.readMem(ic.Pc)%100 != 99 {
		if err := ic.Step(); err != nil {
			return err
		}
	}
	// Execute a last Step() instruction to clean up resources during halt
	return ic.Step()
}

func (ic *Intcode) makeErr(msg string, args ...interface{}) error {
	return fmt.Errorf(msg+" at pc=%d", append(args, ic.Pc)...)
}

// Executes a single instruction and advances the program counter.
func (ic *Intcode) Step() error {
	switch ic.readMem(ic.Pc) % 100 {
	case 1: // Add
		ic.setMem(3, ic.paramVal(1)+ic.paramVal(2))
		ic.Pc += 4
	case 2: // Multiply
		ic.setMem(3, ic.paramVal(1)*ic.paramVal(2))
		ic.Pc += 4
	case 3: // Input
		var inputVal int64
		if ic.CurrentInput < len(ic.Inputs) {
			inputVal = ic.Inputs[ic.CurrentInput]
			ic.CurrentInput++
		} else if ic.InputChan != nil {
			var ok bool
			inputVal, ok = <-ic.InputChan
			if !ok {
				return fmt.Errorf("not enough inputs")
			}
		} else {
			return fmt.Errorf("not enough inputs")
		}
		ic.setMem(1, inputVal)
		ic.Pc += 2
	case 4: // Output
		outputValue := ic.paramVal(1)
		ic.Outputs = append(ic.Outputs, outputValue)
		if ic.OutputChan != nil {
			ic.OutputChan <- outputValue
		}
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
			ic.setMem(3, 1)
		} else {
			ic.setMem(3, 0)
		}
		ic.Pc += 4
	case 8: // Equals
		if ic.paramVal(1) == ic.paramVal(2) {
			ic.setMem(3, 1)
		} else {
			ic.setMem(3, 0)
		}
		ic.Pc += 4
	case 9: // Set relative offset
		ic.RelativeOffset += ic.paramVal(1)
		ic.Pc += 2
	case 99: // Halt
		if ic.OutputChan != nil {
			close(ic.OutputChan)
		}
		if ic.CurrentInput < len(ic.Inputs) {
			return fmt.Errorf("unused input value")
		}
		// Ignore any extra inputs on input channel; for feedback loops we will get some extra inputs that we ignore.
		// Don't insist on input channel being closed before halting.
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

func (ic *Intcode) InitializeProgram(mem ...int64) {
	ic.originalProgram = append([]int64(nil), mem...)
	ic.Reset()
}
