package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

func onComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken
}

func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

type intcodecpu struct {
	memory []int
	input  []int
	output []int
	pc     int
}

func (cpu *intcodecpu) read(addr int, mode int) int {
	if mode == 0 {
		return cpu.memory[cpu.memory[addr]]
	} else if mode == 1 {
		return cpu.memory[addr]
	}
	panic("invalid addr mode")
}

func (cpu *intcodecpu) write(addr int, value int, mode int) {
	cpu.memory[cpu.memory[addr]] = value
}

func getModes(opcode int) (int, int, int) {
	return (opcode / 100 % 10), (opcode / 1000 % 10), (opcode / 10000 % 10)
}

func (cpu *intcodecpu) execOpcodeAdd(pc int) int {
	m1, m2, m3 := getModes(cpu.memory[pc])
	val := cpu.read(pc+1, m1) + cpu.read(pc+2, m2)
	cpu.write(pc+3, val, m3)
	return pc + 4
}

func (cpu *intcodecpu) execOpcodeMul(pc int) int {
	m1, m2, m3 := getModes(cpu.memory[pc])
	val := cpu.read(pc+1, m1) * cpu.read(pc+2, m2)
	cpu.write(pc+3, val, m3)
	return pc + 4
}

func (cpu *intcodecpu) execOpcodeInput(pc int) int {
	m1, _, _ := getModes(cpu.memory[pc])
	val := 0
	// if no input suspend
	if len(cpu.input) == 0 {
		return pc
	}
	val = cpu.input[0]
	cpu.input = cpu.input[1:]
	cpu.write(pc+1, val, m1)

	return pc + 2
}

func (cpu *intcodecpu) execOpcodeOutput(pc int) int {
	m1, _, _ := getModes(cpu.memory[pc])
	val := cpu.read(pc+1, m1)
	fmt.Println(val)
	cpu.output = append(cpu.output, val)
	return pc + 2
}

func (cpu *intcodecpu) execOpcodeJmpTrue(pc int) int {
	m1, m2, _ := getModes(cpu.memory[pc])
	flag := cpu.read(pc+1, m1)
	if flag != 0 {
		return cpu.read(pc+2, m2)
	}
	return pc + 3
}

func (cpu *intcodecpu) execOpcodeJmpFalse(pc int) int {
	m1, m2, _ := getModes(cpu.memory[pc])
	flag := cpu.read(pc+1, m1)
	if flag == 0 {
		return cpu.read(pc+2, m2)
	}
	return pc + 3
}

func (cpu *intcodecpu) execOpcodeLessThan(pc int) int {
	m1, m2, m3 := getModes(cpu.memory[pc])
	p1 := cpu.read(pc+1, m1)
	p2 := cpu.read(pc+2, m2)
	if p1 < p2 {
		cpu.write(pc+3, 1, m3)
	} else {
		cpu.write(pc+3, 0, m3)
	}
	return pc + 4
}

func (cpu *intcodecpu) execOpcodeEquals(pc int) int {
	m1, m2, m3 := getModes(cpu.memory[pc])
	p1 := cpu.read(pc+1, m1)
	p2 := cpu.read(pc+2, m2)
	if p1 == p2 {
		cpu.write(pc+3, 1, m3)
	} else {
		cpu.write(pc+3, 0, m3)
	}
	return pc + 4
}

func (cpu *intcodecpu) execCodeOne() bool {
	opcode := cpu.memory[cpu.pc] % 100
	switch opcode {
	case 1:
		cpu.pc = cpu.execOpcodeAdd(cpu.pc)
	case 2:
		cpu.pc = cpu.execOpcodeMul(cpu.pc)
	case 3:
		cpu.pc = cpu.execOpcodeInput(cpu.pc)
	case 4:
		cpu.pc = cpu.execOpcodeOutput(cpu.pc)
	case 5:
		cpu.pc = cpu.execOpcodeJmpTrue(cpu.pc)
	case 6:
		cpu.pc = cpu.execOpcodeJmpFalse(cpu.pc)
	case 7:
		cpu.pc = cpu.execOpcodeLessThan(cpu.pc)
	case 8:
		cpu.pc = cpu.execOpcodeEquals(cpu.pc)
	case 99:
		return false
	default:
		panic(fmt.Sprintf("invalid opcode %d", opcode))
	}
	return true
}

func (cpu *intcodecpu) execCode() {
	for cpu.execCodeOne() {
	}
}

func calcAmplitude(mcp []int, phases []int) int {

	output := 0
	for _, phase := range phases {
		cpu := intcodecpu{}
		cpu.memory = append([]int{}, mcp...)
		cpu.input = []int{phase, output}
		cpu.execCode()
		output = cpu.output[0]
	}
	return output
}

func calcAmplitudeFeedback(mcp []int, phases []int) int {
	amps := make([]intcodecpu, len(phases))

	for i, phase := range phases {
		amps[i].memory = append([]int{}, mcp...)
		amps[i].input = []int{phase}
		//amps[i].output = amps[(i+1)%len(phases)].input
	}
	amps[0].input = append(amps[0].input, 0)

	out := 0

	for {
		for i := range amps {
			more := amps[i].execCodeOne()

			if i == len(amps)-1 && !more {
				return out
			}

			if len(amps[i].output) > 0 {
				out = amps[i].output[0]
				amps[(i+1)%len(amps)].input = append(amps[(i+1)%len(amps)].input, out)
				amps[i].output = nil
			}
		}
	}
}

func main() {

	input, err := ioutil.ReadFile("input")
	panicOnErr(err, "Cant Read Inputfile")
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(onComma)

	mcp := []int{}
	for scanner.Scan() {
		valStr := scanner.Text()
		val, err := strconv.Atoi(valStr)
		panicOnErr(err, "Invalid Input: "+valStr)
		mcp = append(mcp, val)
	}

	maxThrust := 0
	Perm([]int{0, 1, 2, 3, 4}, func(a []int) {
		thrust := calcAmplitude(mcp, a)
		if thrust > maxThrust {
			maxThrust = thrust
		}
	})

	fmt.Printf("max thrust 1: %d", maxThrust)

	maxThrust2 := 0
	Perm([]int{5, 6, 7, 8, 9}, func(a []int) {
		thrust := calcAmplitudeFeedback(mcp, a)
		if thrust > maxThrust2 {
			maxThrust2 = thrust
		}
	})

	fmt.Printf("max thrust: %d", maxThrust2)

}
