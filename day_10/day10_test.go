package main

import (
	"testing"

	"github.com/0xcafed00d/assert"
)

func Test1(t *testing.T) {
	assert := assert.Make(t)
	sys := system{}
	sys.fillLines([]string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	})

	assert(calcAngleAndDistance(point{4, 4}, point{4, 0})).Equal(float64(0), float64(4))
	assert(calcAngleAndDistance(point{4, 4}, point{8, 4})).Equal(float64(90), float64(4))
	assert(calcAngleAndDistance(point{4, 4}, point{4, 8})).Equal(float64(180), float64(4))
	assert(calcAngleAndDistance(point{4, 4}, point{0, 4})).Equal(float64(270), float64(4))

	assert(sys.calcVisibleFrom(point{3, 4})).Equal(8)
	assert(sys.calcVisibleFrom(point{4, 4})).Equal(7)
}

func Test2(t *testing.T) {
	assert := assert.Make(t)
	sys := system{}
	sys.fillLines([]string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	})

	assert(sys.calcPart2(point{11, 13}).pos).Equal(point{8, 2})
}
