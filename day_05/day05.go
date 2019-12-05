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

func read(memory []int, addr int, mode int) int {
	if mode == 0 {
		return memory[memory[addr]]
	} else if mode == 1 {
		return memory[addr]
	}
	panic("invalid addr mode")
}

func write(memory []int, addr int, value int, mode int) {
	memory[memory[addr]] = value
}

func getModes(opcode int) (int, int, int) {
	return (opcode / 100 % 10), (opcode / 1000 % 10), (opcode / 10000 % 10)
}

func execOpcodeAdd(memory []int, pc int) int {
	m1, m2, m3 := getModes(memory[pc])
	val := read(memory, pc+1, m1) + read(memory, pc+2, m2)
	write(memory, pc+3, val, m3)
	return pc + 4
}

func execOpcodeMul(memory []int, pc int) int {
	m1, m2, m3 := getModes(memory[pc])
	val := read(memory, pc+1, m1) * read(memory, pc+2, m2)
	write(memory, pc+3, val, m3)
	return pc + 4
}

func execOpcodeInput(memory []int, pc int) int {
	fmt.Print("Input:> ")

	m1, _, _ := getModes(memory[pc])
	var inp string
	fmt.Scanln(&inp)
	val, err := strconv.Atoi(inp)
	panicOnErr(err, "Invalid Input: "+inp)
	write(memory, pc+1, val, m1)

	return pc + 2
}

func execOpcodeOutput(memory []int, pc int) int {
	m1, _, _ := getModes(memory[pc])
	val := read(memory, pc+1, m1)
	fmt.Println(val)
	return pc + 2
}

func execOpcodeJmpTrue(memory []int, pc int) int {
	m1, m2, _ := getModes(memory[pc])
	flag := read(memory, pc+1, m1)
	if flag != 0 {
		return read(memory, pc+2, m2)
	}
	return pc + 3
}

func execOpcodeJmpFalse(memory []int, pc int) int {
	m1, m2, _ := getModes(memory[pc])
	flag := read(memory, pc+1, m1)
	if flag == 0 {
		return read(memory, pc+2, m2)
	}
	return pc + 3
}

func execOpcodeLessThan(memory []int, pc int) int {
	m1, m2, m3 := getModes(memory[pc])
	p1 := read(memory, pc+1, m1)
	p2 := read(memory, pc+2, m2)
	if p1 < p2 {
		write(memory, pc+3, 1, m3)
	} else {
		write(memory, pc+3, 0, m3)
	}
	return pc + 4
}

func execOpcodeEquals(memory []int, pc int) int {
	m1, m2, m3 := getModes(memory[pc])
	p1 := read(memory, pc+1, m1)
	p2 := read(memory, pc+2, m2)
	if p1 == p2 {
		write(memory, pc+3, 1, m3)
	} else {
		write(memory, pc+3, 0, m3)
	}
	return pc + 4
}

func execCode(memory []int) []int {
	pc := 0
	for {
		opcode := memory[pc] % 100
		switch opcode {
		case 1:
			pc = execOpcodeAdd(memory, pc)
		case 2:
			pc = execOpcodeMul(memory, pc)
		case 3:
			pc = execOpcodeInput(memory, pc)
		case 4:
			pc = execOpcodeOutput(memory, pc)
		case 5:
			pc = execOpcodeJmpTrue(memory, pc)
		case 6:
			pc = execOpcodeJmpFalse(memory, pc)
		case 7:
			pc = execOpcodeLessThan(memory, pc)
		case 8:
			pc = execOpcodeEquals(memory, pc)
		case 99:
			return memory
		default:
			panic(fmt.Sprintf("invalid opcode %d", opcode))
		}
	}
}

func main() {

	input, err := ioutil.ReadFile("input")
	panicOnErr(err, "Cant Read Inputfile")
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(onComma)

	memory := []int{}

	for scanner.Scan() {
		valStr := scanner.Text()
		val, err := strconv.Atoi(valStr)
		panicOnErr(err, "Invalid Input: "+valStr)
		memory = append(memory, val)
	}

	execCode(memory)
}
