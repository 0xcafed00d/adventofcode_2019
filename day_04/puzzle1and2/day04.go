package main

import (
	"fmt"
	"math"
	"os"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

func getDigit(val int, digit int) int {
	return (val / int(math.Pow10(digit))) % 10
}

func checkPW_part1(pw int) bool {
	pair := false

	lastdigit := 11
	curdigit := 0
	for n := 0; n < 6; n++ {
		curdigit = getDigit(pw, n)
		if lastdigit == curdigit {
			pair = true
		}
		if lastdigit < curdigit {
			return false
		}
		lastdigit = curdigit
	}

	return pair
}

func checkPW_part2(pw int) bool {
	var digitCount [10]int

	lastdigit := 11
	curdigit := 0
	for n := 0; n < 6; n++ {
		curdigit = getDigit(pw, n)
		if lastdigit < curdigit {
			return false
		}
		lastdigit = curdigit
		digitCount[curdigit]++
	}

	for _, cnt := range digitCount {
		if cnt == 2 {
			return true
		}
	}

	return false
}

func main() {
	count1 := 0
	count2 := 0

	for n := 152085; n <= 670283; n++ {
		if checkPW_part1(n) {
			count1++
		}
		if checkPW_part2(n) {
			count2++
		}
	}
	fmt.Println(count1, count2)
}
