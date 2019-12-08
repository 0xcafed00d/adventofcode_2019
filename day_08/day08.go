package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

type layer [25][6]int

func main() {

	layers := []layer{}

	input, err := ioutil.ReadFile("input")
	panicOnErr(err, "Cant Read Input file")

	for i := 0; len(input) > 0; i++ {
		fmt.Println(len(layers), " layers")

		layers = append(layers, layer{})
		for y := 0; y < 6; y++ {
			for x := 0; x < 25; x++ {
				layers[i][x][y] = int(input[0]) - int('0')
				if len(input) > 1 {
					input = input[1:]
				} else {
					input = nil
				}
			}
		}
	}

	minZeros := math.MaxInt32
	minLayer := 0
	min1s := 0
	min2s := 0

	for i := range layers {
		counts := [3]int{0, 0, 0}
		for y := 0; y < 6; y++ {
			for x := 0; x < 25; x++ {
				val := layers[i][x][y]
				counts[val]++
			}
		}

		if counts[0] < minZeros {
			minLayer = i
			minZeros = counts[0]
			min1s = counts[1]
			min2s = counts[2]
		}
	}
	fmt.Println(len(layers), " layers")

	fmt.Printf("min zeros on on layer %d (%d zeros) 1s * 2s = %d\n", minLayer, minZeros, min1s*min2s)

	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			pixel := 0
			for i := len(layers) - 1; i >= 0; i-- {
				p := layers[i][x][y]
				if p != 2 {
					pixel = p
				}
			}
			if pixel == 1 {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
