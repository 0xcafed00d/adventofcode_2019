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

func calcFuelForMass(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func calcFuel(mass int) int {
	fuel := 0
	for {
		f := calcFuelForMass(mass)
		if f == 0 {
			break
		}
		fuel += f
		mass = f
	}
	return fuel
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
