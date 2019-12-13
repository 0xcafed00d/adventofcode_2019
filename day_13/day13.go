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

type intcodecpu struct {
	memory   map[int]int64
	input    []int64
	output   []int64
	pc       int
	relbase  int
	inputreq bool
}

func (cpu *intcodecpu) read(addr int, mode int) int64 {

	if mode == 0 {
		return cpu.memory[int(cpu.memory[addr])]
	} else if mode == 1 {
		return cpu.memory[addr]
	} else if mode == 2 {
		return cpu.memory[cpu.relbase+int(cpu.memory[addr])]
	}
	panic("invalid addr mode")
}

func (cpu *intcodecpu) write(addr int, value int64, mode int) {
	if mode == 0 {
		cpu.memory[int(cpu.memory[addr])] = value
	} else if mode == 2 {
		cpu.memory[cpu.relbase+int(cpu.memory[addr])] = value
	}
}

func getModes(opcode int64) (int, int, int) {
	return int(opcode / 100 % 10), int(opcode / 1000 % 10), int(opcode / 10000 % 10)
}

func (cpu *intcodecpu) setProgram(p []int64) {
	cpu.memory = make(map[int]int64)

	for addr, val := range p {
		cpu.memory[addr] = val
	}
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
	val := int64(0)
	// if no input suspend
	if len(cpu.input) == 0 {
		cpu.inputreq = true
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
	//fmt.Println(val)
	cpu.output = append(cpu.output, val)
	return pc + 2
}

func (cpu *intcodecpu) execOpcodeJmpTrue(pc int) int {
	m1, m2, _ := getModes(cpu.memory[pc])
	flag := cpu.read(pc+1, m1)
	if flag != 0 {
		return int(cpu.read(pc+2, m2))
	}
	return pc + 3
}

func (cpu *intcodecpu) execOpcodeJmpFalse(pc int) int {
	m1, m2, _ := getModes(cpu.memory[pc])
	flag := cpu.read(pc+1, m1)
	if flag == 0 {
		return int(cpu.read(pc+2, m2))
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

func (cpu *intcodecpu) execOpcodeSetRel(pc int) int {
	m1, _, _ := getModes(cpu.memory[pc])
	p1 := cpu.read(pc+1, m1)
	cpu.relbase += int(p1)
	return pc + 2
}

func (cpu *intcodecpu) execCodeOne() bool {
	cpu.inputreq = false
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
	case 9:
		cpu.pc = cpu.execOpcodeSetRel(cpu.pc)
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

type point struct {
	x int64
	y int64
}

var dirs = []point{
	point{0, -1}, // up
	point{1, 0},  // right
	point{0, 1},  // down
	point{-1, 0}, // left
}

func runGame(cpu *intcodecpu, screen map[point]int64) {

	for cpu.execCodeOne() {
		if cpu.inputreq {
			panic("input requested")
		}

		if len(cpu.output) > 2 {
			screen[point{cpu.output[0], cpu.output[1]}] = cpu.output[2]
			cpu.output = cpu.output[3:]
		}
	}
}

func runGame2(cpu *intcodecpu, screen map[point]int64) int64 {

	ballx := int64(0)
	paddlex := int64(0)
	score := int64(0)

	for cpu.execCodeOne() {
		if len(cpu.output) > 2 {
			screen[point{cpu.output[0], cpu.output[1]}] = cpu.output[2]

			if cpu.output[0] == -1 && cpu.output[1] == 0 {
				score = cpu.output[2]
			} else {
				if cpu.output[2] == 4 { //ball
					ballx = cpu.output[0]
				}
				if cpu.output[2] == 3 { //paddle
					paddlex = cpu.output[0]
				}
			}
			cpu.output = cpu.output[3:]
		}
		if cpu.inputreq {
			if ballx == paddlex {
				cpu.input = []int64{0}
			} else if ballx < paddlex {
				cpu.input = []int64{-1}
			} else if ballx > paddlex {
				cpu.input = []int64{1}
			}
		}
	}
	return score
}

func main() {

	input, err := ioutil.ReadFile("input")
	panicOnErr(err, "Cant Read Inputfile")
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(onComma)

	mcp := []int64{}
	for scanner.Scan() {
		valStr := scanner.Text()
		val, err := strconv.ParseInt(valStr, 10, 64)
		panicOnErr(err, "Invalid Input: "+valStr)
		mcp = append(mcp, val)
	}

	cpu := intcodecpu{}
	cpu.setProgram(mcp)
	screen := make(map[point]int64)
	runGame(&cpu, screen)

	count := 0
	for _, v := range screen {
		if v == 2 {
			count++
		}
	}

	fmt.Println(count)

	cpu = intcodecpu{}
	cpu.setProgram(mcp)
	screen = make(map[point]int64)
	cpu.memory[0] = 2
	score := runGame2(&cpu, screen)

	fmt.Println(score)
}
