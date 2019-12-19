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

func pattern(digit int) func() int {
	src := []int{0, 1, 0, -1}

	count := 1
	index := 0
	return func() int {
		index = (count / (digit + 1)) & 3
		val := src[index]
		count++
		return val
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func doPhase(inp []int) (output []int) {
	output = make([]int, len(inp), len(inp))
	for i := range inp {
		fmt.Println(i)
		pgen := pattern(i)
		tot := 0
		for _, v := range inp {
			tot += v * pgen()
		}
		output[i] = abs(tot) % 10
	}
	return
}

func doPhase2(inp []int, offset int) (output []int) {
	sz := len(inp)
	output = make([]int, len(inp), len(inp))
	for oi := offset; oi < sz; oi++ {
		if oi&0xff == 0 {
			fmt.Println(oi)
		}
		tot := 0
		for ii := oi; ii < sz; ii++ {
			tot += inp[ii]
		}
		output[oi] = abs(tot) % 10
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		for i := range line {
			n, err := strconv.Atoi(line[i : i+1])
			panicOnErr(err, "parse error")
			input = append(input, n)
		}
	}

	part1 := append([]int{}, input...)
	for n := 0; n < 100; n++ {
		part1 = doPhase(part1)
	}
	fmt.Println(part1[0:8])

	offset := 0
	for n := 0; n < 7; n++ {
		offset *= 10
		offset += input[n]
	}

	part2 := []int{}
	for n := 0; n < 10000; n++ {
		part2 = append(part2, input...)
	}

	fmt.Println(offset)
	fmt.Println(len(part2))

	for n := 0; n < 100; n++ {
		fmt.Print("phase: ", n)
		part2 = doPhase2(part2, offset)
	}
	fmt.Println("part1  ", part1[0:8])
	fmt.Println("part2   ", part2[offset:offset+10])
}
