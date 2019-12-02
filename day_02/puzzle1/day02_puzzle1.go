package main

import (
	"bufio"
	"fmt"
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

func execOpcodeAdd(memory []int, pc int) int {
	memory[memory[pc+3]] = memory[memory[pc+1]] + memory[memory[pc+2]]
	return pc + 4
}

func execOpcodeMul(memory []int, pc int) int {
	memory[memory[pc+3]] = memory[memory[pc+1]] * memory[memory[pc+2]]
	return pc + 4
}

func execCode(memory []int) []int {
	pc := 0
	for {
		switch memory[pc] {
		case 1:
			pc = execOpcodeAdd(memory, pc)
		case 2:
			pc = execOpcodeMul(memory, pc)
		case 99:
			return memory
		default:
			panic("invalid opcode")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(onComma)

	memory := []int{}

	for scanner.Scan() {
		valStr := scanner.Text()
		val, err := strconv.Atoi(valStr)
		panicOnErr(err, "Invalid Input: "+valStr)
		memory = append(memory, val)
	}

	fmt.Println(memory)

	memory[1] = 12
	memory[2] = 2

	execCode(memory)
	fmt.Println(memory)

}
