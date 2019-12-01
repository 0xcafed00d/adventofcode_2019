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

func calcFuel(mass int) int {
	return (mass / 3) - 2
}

func main() {
	fuelTotal := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		massStr := scanner.Text()
		mass, err := strconv.Atoi(massStr)
		panicOnErr(err, "Invalid Input: "+massStr)

		fuelTotal += calcFuel(mass)
	}

	fmt.Printf("Total Fuel Required: %d ", fuelTotal)
}
