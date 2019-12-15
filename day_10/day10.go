package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func panicOnErr(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		panic(err)
	}
}

type point struct {
	x int
	y int
}

type body struct {
	present      bool
	visibleCount int
	isVisible    int
}

type system [100][100]body

func (s *system) fillLine(l int, line string) {
	for i, v := range line {
		if v == '#' {
			(*s)[l][i].present = true
		}
	}
}

func (s *system) fillLines(lines []string) {
	for l, v := range lines {
		s.fillLine(l, v)
	}
}

func (s *system) getBody(p point) *body {
	return &((*s)[p.y][p.x])
}

func calcAngleAndDistance(from, to point) (angle, dist float64) {
	dx := float64(from.x - to.x)
	dy := float64(from.y - to.y)
	angle = math.Atan2(-dx, dy) * (180 / math.Pi)
	angle = math.Mod(angle+360, 360)
	dist = math.Sqrt(dx*dx + dy*dy)
	return
}

func (s *system) calcVisibleFrom(from point) int {
	angles := make(map[float64]bool)

	for x := 0; x < len((*s)[0]); x++ {
		for y := 0; y < len(*s); y++ {
			to := point{x, y}
			if from != to && s.getBody(to).present {
				a, _ := calcAngleAndDistance(from, to)
				angles[a] = true
			}
		}
	}
	return len(angles)
}

func (s *system) calcPart1() (int, point) {
	maxVis := 0
	maxVisPnt := point{}

	for x := 0; x < len((*s)[0]); x++ {
		for y := 0; y < len(*s); y++ {
			p := point{x, y}
			if s.getBody(p).present {
				vis := s.calcVisibleFrom(p)
				if vis > maxVis {
					maxVis = vis
					maxVisPnt = p
				}
			}
		}
	}
	return maxVis, maxVisPnt
}

type target struct {
	pos       point
	angle     float64
	dist      float64
	destroyed bool
}

func (s *system) calcPart2(src point) target {

	targets := []target{}

	for x := 0; x < len((*s)[0]); x++ {
		for y := 0; y < len(*s); y++ {
			p := point{x, y}

			if p != src && s.getBody(p).present {
				a, d := calcAngleAndDistance(src, p)
				targets = append(targets, target{p, a, d, false})
			}
		}
	}

	sort.Slice(targets, func(i, j int) bool {
		if targets[i].angle == targets[j].angle {
			return targets[i].dist < targets[j].dist
		}
		return targets[i].angle < targets[j].angle
	})
	/*
		for i := range targets {
			fmt.Println(targets[i])
		}
	*/

	lastangle := float64(-1)
	destroyed := []target{}
	i := 0
	for {
		t := &targets[i]
		if !t.destroyed && t.angle != lastangle {
			t.destroyed = true
			destroyed = append(destroyed, *t)
			fmt.Println(*t)
			lastangle = t.angle
		}

		if len(destroyed) == len(targets) {
			break
		}

		i++
		if i == len(targets) {
			i = 0
			lastangle = -1
		}
	}

	return destroyed[199]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sys := system{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		sys.fillLine(i, line)
		i++
	}

	count, pnt := sys.calcPart1()
	fmt.Println(count, pnt)

	t := sys.calcPart2(pnt)
	fmt.Println(t)
}
