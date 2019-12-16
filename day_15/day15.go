package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
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
	x int
	y int
}

func (a point) add(b point) point {
	return point{a.x + b.x, a.y + b.y}
}

var dirs = []point{
	point{0, -1}, // up
	point{1, 0},  // right
	point{0, 1},  // down
	point{-1, 0}, // left
}

var dirinput = []int64{
	1, // up    / north
	4, // right / east
	2, // down  / south
	3, // left  / west
}

func shortestDistance(maze map[point]rune, from point, to point, count int) (int, bool) {
	fmt.Println(from)
	if from == to {
		return count, true
	}
	if maze[from] != ' ' {
		return 0, false
	}
	maze[from] = '.'

	for dir := 0; dir < 4; dir++ {
		n, ok := shortestDistance(maze, from.add(dirs[dir]), to, count+1)
		if ok {
			return n, ok
		}
	}

	return 0, false
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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
	maze := make(map[point]rune)

	start := point{25, 25}
	end := point{}

	pos := start
	dir := 0

	maze[pos] = 'S'
	for cpu.execCodeOne() {
		if cpu.inputreq {
			if maze[pos.add(dirs[dir])] != 0 {
				dir = (dir + rand.Intn(4)) % 4
			}
			cpu.input = append(cpu.input, dirinput[dir])
		}

		if len(cpu.output) > 0 {
			out := cpu.output[0]
			cpu.output = cpu.output[1:]
			if out == 0 {
				maze[pos.add(dirs[dir])] = '#'
			} else if out == 1 {
				pos = pos.add(dirs[dir])
				maze[pos] = ' '
			} else if out == 2 {
				pos = pos.add(dirs[dir])
				maze[pos] = 'X'
				end = pos
				break
			}
		}
	}

	//maze[point{25, 25}] = 'S'
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			p := point{x, y}
			v := maze[p]
			if v == 0 {
				maze[p] = '#'
			}
			fmt.Printf("%c", maze[p])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println(shortestDistance(maze, start, end, 0))
}
