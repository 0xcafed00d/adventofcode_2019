package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

func onSep(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' || data[i] < ' ' {
			return i + 1, data[:i], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken
}

type point struct {
	x int
	y int
}

var zeroPos = point{0, 0}

var dirs = map[byte]point{
	'U': {0, 1},
	'D': {0, -1},
	'L': {-1, 0},
	'R': {1, 0},
}

func calcDist(wires []string) (int, int) {
	circuit := make(map[point][2]int)
	mdist := math.MaxInt32
	msteps := math.MaxInt32

	for i, wire := range wires {
		scanner := bufio.NewScanner(strings.NewReader(wire))
		scanner.Split(onSep)
		pos := zeroPos
		steps := 0
		for scanner.Scan() {
			move := scanner.Text()
			dir := dirs[move[0]]
			dist, err := strconv.Atoi(move[1:])
			panicOnErr(err, move[1:])

			for n := 0; n < dist; n++ {
				a := circuit[pos]
				a[i] = steps
				circuit[pos] = a

				if circuit[pos][0] > 0 && circuit[pos][1] > 0 {
					mdist = Min(Abs(pos.x)+Abs(pos.y), mdist)
					msteps = Min(circuit[pos][0]+circuit[pos][1], msteps)
				}
				pos.x += dir.x
				pos.y += dir.y
				steps++
			}
		}
	}
	return mdist, msteps
}

func main() {

	wires := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		wire := scanner.Text()
		wires = append(wires, wire)
	}

	fmt.Println(calcDist(wires))
}
