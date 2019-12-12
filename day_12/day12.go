package main

import (
	"fmt"
	"os"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

type point struct {
	x, y, z, dx, dy, dz int
}

func calcSign(p1, p2 int) int {
	if p1 > p2 {
		return -1
	}
	if p1 < p2 {
		return 1
	}
	return 0
}

func calcVel(state []point) {
	for i1 := range state {
		for i2 := range state {
			state[i2].dx += calcSign(state[i2].x, state[i1].x)
			state[i2].dy += calcSign(state[i2].y, state[i1].y)
			state[i2].dz += calcSign(state[i2].z, state[i1].z)
		}
	}
}

func applyVel(state []point) {
	for i := range state {
		state[i].x += state[i].dx
		state[i].y += state[i].dy
		state[i].z += state[i].dz
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calcEnergy(state []point) int {
	e := 0
	for i := range state {
		pe := abs(state[i].x) + abs(state[i].y) + abs(state[i].z)
		ke := abs(state[i].dx) + abs(state[i].dy) + abs(state[i].dz)
		e += (pe * ke)
	}
	return e
}

func getInput() []point {
	return []point{
		point{x: -10, y: -13, z: 7},
		point{x: 1, y: 2, z: 1},
		point{x: -15, y: -3, z: 13},
		point{x: 3, y: 7, z: -4},
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func main() {

	input := getInput()

	for i := 0; i < 1000; i++ {
		calcVel(input)
		applyVel(input)
	}

	fmt.Println(calcEnergy(input))

	input = getInput()
	initial := getInput()

	count := int64(0)

	counts := [3]int64{}
	// find cycle period in each axis
	for {
		calcVel(input)
		applyVel(input)

		count++

		if counts[0] == 0 {
			match := true
			for i := range input {
				if input[i].x != initial[i].x || input[i].dx != initial[i].dx {
					match = false
				}
			}
			if match {
				counts[0] = count
			}
		}

		if counts[1] == 0 {
			match := true
			for i := range input {
				if input[i].y != initial[i].y || input[i].dy != initial[i].dy {
					match = false
				}
			}
			if match {
				counts[1] = count
			}
		}

		if counts[2] == 0 {
			match := true
			for i := range input {
				if input[i].z != initial[i].z || input[i].dz != initial[i].dz {
					match = false
				}
			}
			if match {
				counts[2] = count
			}
		}

		if counts[0] > 0 && counts[1] > 0 && counts[2] > 0 {
			break
		}

	}

	fmt.Println(counts[0], counts[1], counts[2])

	// find lowest common multiple of the periods of esch axis
	fmt.Println(LCM(counts[0], counts[1], counts[2]))

}
